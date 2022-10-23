package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"GoDBProject/internal/application/models"
	"GoDBProject/pkg/storage/postgres"
)

// Storage это структура, которая создаёт пулл соединений к постгресу
type Storage struct {
	pl *pgxpool.Pool
}

func NewStorage(conf postgres.Config) (*Storage, error) {
	pool, err := conf.Pool()
	if err != nil {
		return nil, fmt.Errorf("создание пула соединений постгреса: %w", err)
	}
	return &Storage{
		pl: pool,
	}, nil
}

func (s *Storage) Close() {
	s.pl.Close()
}

func (s *Storage) InsertPTS(ctx context.Context, service int, product int, count int) error {
	const query = `CALL insert_service_for_relise($1, $2, $3)`
	_, err := s.pl.Exec(ctx, query, &product, &service, &count)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg: %w", err)
	}
	return nil
}

func (s *Storage) InsertGroupServices(ctx context.Context, services models.GroupServices) error {
	const query = `CALL insert_group_services($1, $2, $3)`
	_, err := s.pl.Exec(ctx, query, &services.Name, &services.Code, &services.Hidden)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg: %w", err)
	}
	return nil
}
func (s *Storage) InsertProduct(ctx context.Context, product models.Product) error {
	const query = `CALL insert_service_for_relise($1, $2, $3)`
	_, err := s.pl.Exec(ctx, query, &product.Code, &product.Name, &product.MeasureUnit, &product.Basecost, &product.Hidden)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg: %w", err)
	}
	return nil
}
func (s *Storage) InsertService(ctx context.Context, service models.Service) error {
	const query = `CALL insert_services($1, $2, $3, $4, $5)`
	_, err := s.pl.Exec(ctx, query, &service.Code, &service.Name, &service.GroupId, &service.Basecost, &service.Hidden)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg: %w", err)
	}
	return nil
}
func (s *Storage) InsertMU(ctx context.Context, unit models.MeasureUnit) error {
	const query = `CALL insert_measure_unit($1, $2)`
	_, err := s.pl.Exec(ctx, query, &unit.Name, &unit.ShortName)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg: %w", err)
	}
	return nil
}

// GetAll метод возращающий прайс листы из БД
func (s *Storage) GetAll(ctx context.Context) ([]models.Price, error) {
	const query = `SELECT ps."name" as service_name,
       sum(count_of_prod*pp.basecost+ps.basecost) as price,
       sum(ps.basecost) as serviceprice,
       sum(pp.basecost) as materialprice
FROM public.prod_to_service
         join public.products pp ON pp.id = prod_to_service.prod_id
         join public.services ps ON ps.id = prod_to_service.service_id
group by service_name;`
	var (
		buf       models.Price
		resultMap []models.Price
	)
	rows, err := s.pl.Query(ctx, query)
	if err != nil {
		return resultMap, fmt.Errorf("ошибка при запросе на получение блока по GUID'у %w", err)
	}
	for rows.Next() {
		// err = rows.Scan(&bufMap.Id, &bufMap.Name, &bufMap.ShortName)
		err = rows.Scan(&buf.ServiceName, &buf.Price, &buf.ServicePrice, &buf.MaterialPrice)
		if err != nil {
			return resultMap, fmt.Errorf("error on scanning answers: %w", err)
		}

		resultMap = append(resultMap, buf)
	}
	return resultMap, nil
}

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

func (s *Storage) InsertPTS(ctx context.Context, product int, service int, count int) error {
	const query = `CALL insert_service_for_relise($1, $2, $3)`
	_, err := s.pl.Exec(ctx, query, &product, &service, &count)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg {%d%d%d}: %w", product, service, count, err)
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
       sum(count_of_prod*pp.basecost)+ps.basecost as price,
       ps.basecost as serviceprice,
       sum(count_of_prod*pp.basecost) as materialprice
FROM public.prod_to_service
         join public.products pp ON pp.id = prod_to_service.prod_id
         join public.services ps ON ps.id = prod_to_service.service_id
where ps.hidden = false
group by service_name, ps.basecost;`
	var (
		buf       models.Price
		resultMap []models.Price
	)
	rows, err := s.pl.Query(ctx, query)
	if err != nil {
		return resultMap, fmt.Errorf("ошибка при выполнении запроса GetAll %w", err)
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

// GetMaterialByIdService метод возращающий материальный состав услуги из БД
func (s *Storage) GetMaterialByIdService(ctx context.Context, id int) ([]models.MaterialToService, error) {
	const query = `SELECT ps."name" as service_name,
       pp."name" as product_name,
       count_of_prod as count_prod,
		(select mu."short_name" from public.measure_unit as mu where mu.id = pp.measure_id) as Measure,
       pp.basecost as materialprice
FROM public.prod_to_service
         join public.products pp ON pp.id = prod_to_service.prod_id
         join public.services ps ON ps.id = prod_to_service.service_id
--          join public.group_services pgs on pgs.id = ps.group_id
--          join public.measure_unit pmu on pmu.id = pp.measure_id
where ps.id = $1
;`
	var (
		buf       models.MaterialToService
		resultMap []models.MaterialToService
	)
	rows, err := s.pl.Query(ctx, query, &id)
	if err != nil {
		return resultMap, fmt.Errorf("ошибка при выполнении запроса GetMaterialByIdService %w", err)
	}
	for rows.Next() {
		// err = rows.Scan(&bufMap.Id, &bufMap.Name, &bufMap.ShortName)
		err = rows.Scan(&buf.ServiceName, &buf.ProductName, &buf.CountProduct, &buf.MeasureUnit, &buf.MaterialPrice)
		if err != nil {
			return resultMap, fmt.Errorf("error on scanning answers: %w", err)
		}

		resultMap = append(resultMap, buf)
	}
	return resultMap, nil
}

func (s *Storage) UpdateProductPrice(ctx context.Context, productId int, price float64) error {
	const query = `CALL update_product_price($1, $2)`
	_, err := s.pl.Exec(ctx, query, &productId, &price)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg {%d%f}: %w", productId, price, err)
	}
	return nil
}

func (s *Storage) UpdateServicePrice(ctx context.Context, serviceId int, price float64) error {
	const query = `CALL update_services_price($1, $2)`
	_, err := s.pl.Exec(ctx, query, &serviceId, &price)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg {%d%f}: %w", serviceId, price, err)
	}
	return nil
}

func (s *Storage) UpdateProductHidden(ctx context.Context, productId int, hidden bool) error {
	const query = `CALL update_product_hidden($1, $2)`
	_, err := s.pl.Exec(ctx, query, &productId, &hidden)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg {%d%t}: %w", productId, hidden, err)
	}
	return nil
}

func (s *Storage) UpdateServiceHidden(ctx context.Context, serviceId int, hidden bool) error {
	const query = `CALL update_service_hidden($1, $2)`
	_, err := s.pl.Exec(ctx, query, &serviceId, &hidden)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg {%d%t}: %w", serviceId, hidden, err)
	}
	return nil
}

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
func (s *Storage) InsertProduct(ctx context.Context, service int, product int, count int) error {
	const query = `CALL insert_service_for_relise($1, $2, $3)`
	_, err := s.pl.Exec(ctx, query, &product, &service, &count)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg: %w", err)
	}
	return nil
}
func (s *Storage) InsertService(ctx context.Context, service int, product int, count int) error {
	const query = `CALL insert_service_for_relise($1, $2, $3)`
	_, err := s.pl.Exec(ctx, query, &product, &service, &count)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg: %w", err)
	}
	return nil
}
func (s *Storage) InsertMU(ctx context.Context, service int, product int, count int) error {
	const query = `CALL insert_service_for_relise($1, $2, $3)`
	_, err := s.pl.Exec(ctx, query, &product, &service, &count)
	if err != nil {
		return fmt.Errorf("ошибка в выполнении запроса к pg: %w", err)
	}
	return nil
}

func (s *Storage) GetAll(ctx context.Context) ([]models.MeasureUnit, error) {
	const query = `SELECT id, name, short_name
	FROM public.measure_unit`
	var (
		resultMap []models.MeasureUnit
		bufMap    models.MeasureUnit
	)
	rows, err := s.pl.Query(ctx, query)
	if err != nil {
		return resultMap, fmt.Errorf("ошибка при запросе на получение блока по GUID'у %w", err)
	}
	for rows.Next() {
		err = rows.Scan(&bufMap.Id, &bufMap.Name, &bufMap.ShortName)
		if err != nil {
			return resultMap, fmt.Errorf("error on scanning answers: %w", err)
		}
		resultMap = append(resultMap, bufMap)
	}
	return resultMap, nil
}

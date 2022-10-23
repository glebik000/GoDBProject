CREATE PROCEDURE insert_services(code varchar, name varchar,
                                 group_id integer, basecost float, hidden bool )
    LANGUAGE SQL
    AS $$
INSERT INTO public.services(code, name, group_id, basecost, hidden)
   VALUES ($1, $2, $3, $4, $5);
$$;

-- CALL insert_services('В 01', 'Динамическое наблюдение', 1, 600, true);

CREATE PROCEDURE insert_products(code varchar, name varchar, measure_id integer,
                                 basecost numeric, hidden bool)
    LANGUAGE SQL
    AS $$
INSERT INTO public.products(code, name, measure_id, basecost, hidden)
   VALUES ($1, $2, $3, $4, $5);
$$;

-- CALL insert_products('РМ 01', 'Набор одноразовых инструментов', 2, 150, true);

CREATE PROCEDURE insert_group_services(code varchar, name varchar, hidden bool)
    LANGUAGE SQL
    AS $$
INSERT INTO public.group_services(code, name, hidden)
    VALUES ($1, $2, $3);
$$;

-- CALL insert_data(1, 'В', 'Врачебные манипуляции', true);

CREATE PROCEDURE insert_measure_unit(name varchar, short_name varchar)
    LANGUAGE SQL
    AS $$
INSERT INTO public.measure_unit(name, short_name)
    VALUES ($1, $2);
$$;

-- CALL insert_measure_unit('Штука', 'Шт.');

CREATE PROCEDURE insert_service_for_relise(prod_id integer, service_id integer, count_of_prod integer)
    LANGUAGE SQL
    AS $$
INSERT INTO public.prod_to_service(prod_id, service_id, count_of_prod)
VALUES ($1, $2, $3);
$$;

-- CALL insert_service_for_relise(3, 4, 2);

CREATE PROCEDURE update_services_price(id_service integer, val numeric)
    LANGUAGE SQL
    AS $$
UPDATE public.services
SET basecost = val
WHERE id = id_service;
$$;

-- CALL update_services_price(1, 20.0);

CREATE PROCEDURE update_product_price(id_product integer, val numeric)
    LANGUAGE SQL
    AS $$
UPDATE public.products
SET basecost = val
WHERE id = id_product;
$$;

-- CALL update_product_price(1, 20.0);
-- Получение прайс-листа с услугами
SELECT ps."name" as service_name,
       sum(count_of_prod*pp.basecost)+ps.basecost as price,
       ps.basecost as serviceprice,
       sum(count_of_prod*pp.basecost) as materialprice
FROM public.prod_to_service
         join public.products pp ON pp.id = prod_to_service.prod_id
         join public.services ps ON ps.id = prod_to_service.service_id
where ps.hidden = false
group by service_name, ps.basecost;

-- SELECT ps."name" as service_name,
--        sum(count_of_prod*pp.basecost)+ps.basecost as price,
--        ps.basecost as serviceprice,
--        sum(count_of_prod*pp.basecost) as materialprice
-- FROM public.prod_to_service pts,
--   public.services ps,
--   public.products pp,
-- where pts.prod_id = pp.id
--   and pts.service_id = ps.id
--   and ps.hidden = false
--   ;


-- Получение материального состава услуги
SELECT ps."name" as service_name,
       pp."name" as product_name,
       count_of_prod as count_prod,
       pmu."name" as Measure,
       pp.basecost as materialprice,

FROM public.prod_to_service
         join public.products pp ON pp.id = prod_to_service.prod_id
         join public.services ps ON ps.id = prod_to_service.service_id
         join public.group_services pgs on pgs.id = ps.group_id
         join public.measure_unit pmu on pmu.id = pp.measure_id
WHERE ps.id = $1
;

-- Получение материального состава всех услуг
SELECT ps."name" as service_name,
       pp."name" as product_name,
       count_of_prod as count_prod,
       pmu."name" as Measure,
       pp.basecost as materialprice,

FROM public.prod_to_service
         join public.products pp ON pp.id = prod_to_service.prod_id
         join public.services ps ON ps.id = prod_to_service.service_id
         join public.group_services pgs on pgs.id = ps.group_id
         join public.measure_unit pmu on pmu.id = pp.measure_id
;
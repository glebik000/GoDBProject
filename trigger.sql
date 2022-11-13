
CREATE FUNCTION check_hidden() RETURNS trigger AS $check_hidden$
BEGIN
  IF NEW.hidden = TRUE THEN
    UPDATE public.services
    SET hidden = true
    WHERE group_id = new.id;
  END IF;
        
  IF NEW.hidden = FALSE THEN
    UPDATE public.services
    SET hidden = FALSE
    WHERE group_id = new.id;
  END IF;

  RETURN NEW;
END;
$check_hidden$ LANGUAGE plpgsql;


CREATE TRIGGER check_hidden AFTER UPDATE ON public.group_services
  FOR EACH ROW EXECUTE FUNCTION public.check_hidden();

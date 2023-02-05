DROP TABLE if exists "event_settings";
DROP TABLE if exists "event";
DROP TRIGGER if exists create_event_settings_trigger ON "event";
DROP FUNCTION if exists create_event_settings;
DROP TYPE if exists event_type_enum;
DROP TYPE if exists event_visibility_enum;
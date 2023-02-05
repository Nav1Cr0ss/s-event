CREATE TYPE event_type_enum AS ENUM ('online', 'offline', 'mixed');
CREATE TABLE "event"
(
    "id"          serial PRIMARY KEY,
    "author"      varchar(255)    NOT NULL,
    "title"       varchar(255)    NOT NULL,
    "description" varchar(255)    NOT NULL,
    "type"        event_type_enum NOT NULL,
    "created_at"  timestamp       NOT NULL DEFAULT NOW()
);

CREATE TYPE event_visibility_enum AS ENUM ('male', 'female', 'third');

CREATE TABLE "event_settings"
(
    "id"               serial PRIMARY KEY,
    "event_id"         int                     not null references "event" (id) ON DELETE CASCADE,
    "max_participants" int                     NOT NULL default 1,
    "min_participants" int                     NOT NULL default 1,
    "visibility"       event_visibility_enum[] not null default '{}'::event_visibility_enum[]

);


CREATE FUNCTION create_event_settings() RETURNS trigger AS
$$
BEGIN
    insert into event_settings (event_id) values (new.id);
    RETURN NEW;
end;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_event_settings_trigger
    AFTER INSERT
    ON event
    FOR EACH ROW
EXECUTE PROCEDURE create_event_settings();
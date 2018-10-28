create table if not exists main
(
  accountid serial                                              not null
    constraint main_pkey
    primary key,
  name      varchar(50) default 'john doe' :: character varying not null,
  realm     varchar(50)                                         not null,
  region    varchar(5) default 'eu' :: character varying        not null
);

alter table main
  owner to admin;

create unique index if not exists main_accountid_uindex
  on main (accountid);

create table if not exists guild
(
  name        varchar(100) not null,
  realm       varchar(50)  not null,
  region      varchar(5)   not null,
  officerrank integer      not null,
  raiderrank  integer      not null,
  trialrank   integer      not null,
  guildid     serial       not null
    constraint guild_pk
    primary key,
  constraint "Guild-ID"
  unique (name, realm, region)
);

alter table guild
  owner to admin;

create index if not exists guild_name_index
  on guild (name);

create index if not exists guild_realm_index
  on guild (realm);

create unique index if not exists guild_guildid_uindex
  on guild (guildid);

create table if not exists raidnight
(
  dayoftheweek integer not null,
  guildid      integer not null
    constraint raidnight_guild_guildid_fk
    references guild,
  id           serial  not null
    constraint raidnight_pk
    primary key,
  duration     bigint  not null,
  start        bigint  not null
);

alter table raidnight
  owner to admin;

create index if not exists raidnight_guildid_index
  on raidnight (guildid);

create unique index if not exists raidnight_id_uindex
  on raidnight (id);

create table if not exists weakauras
(
  guildid integer      not null
    constraint weakauras_guild_guildid_fk
    references guild,
  name    varchar(150) not null,
  link    varchar(250) not null,
  import  bytea,
  id      serial       not null
    constraint weakauras_pk
    primary key
);

alter table weakauras
  owner to admin;

create index if not exists addons_guildid_index
  on weakauras (guildid);

create unique index if not exists weakauras_id_uindex
  on weakauras (id);

create table if not exists addons
(
  name       varchar(50)  not null,
  twitchlink varchar(250) not null,
  guildid    integer      not null
    constraint addons_guild_guildid_fk
    references guild,
  id         serial       not null
    constraint addons_pk
    primary key
);

alter table addons
  owner to admin;

create unique index if not exists addons_id_uindex
  on addons (id);


INSERT INTO character_typ (character_code, name, created_at, updated_at) 
VALUES (1, 'Wizard', '2019-10-01 15:36:38'::timestamp, '2019-10-01 15:36:38'::timestamp),
        (2, 'Elf', '2019-10-01 15:36:38'::timestamp, '2019-10-01 15:36:38'::timestamp),
        (3, 'Hobbit', '2019-10-01 15:36:38'::timestamp, '2019-10-01 15:36:38'::timestamp);

INSERT INTO character (id, name, character_code, character_power, character_value, created_at, updated_at)
VALUES ('967d5bb5-3a7a-4d5e-8a6c-febc8c5b3f14', 'Gandalf', 1, 100, 150, '2019-10-01 15:36:38'::timestamp, '2019-10-01 15:36:38'::timestamp),
       ('c809bf15-bc2c-4621-bb96-70af96fd5d68', 'Legolas', 2, 60, 68, '2019-10-02 11:16:12'::timestamp, '2019-10-02 11:16:12'::timestamp),
       ('2367710a-d4fb-49f5-8860-557b337386de', 'Frodo', 3, 10, 20, '2019-10-05 05:21:11'::timestamp, '2019-10-05 05:21:11'::timestamp);
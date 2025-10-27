#Desenvolvimento Web com Go
#Rotas da aplicacao

|POST | /note/create | noteCreate | Cria uma aplicacao 



##Modelo do banco de dados

Campo      |   Tipo    |  Constraint
ID         | BIGSERIAL |    PK, NOT NULL
TITLE      | TEXT      |    NOT NULL
CONTENT    | TEXT      |
COLOR      | TEXT      |    NOT NULL
CREATED_AT | TIMESTAMP |
UPDATED_AT | TIMESTAMP |
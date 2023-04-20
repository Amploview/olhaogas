CREATE TABLE IF NOT EXISTS `area` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `descricao` text(80) NOT NULL,
  `ts` text NOT NULL DEFAULT CURRENT_TIMESTAMP
) ;

INSERT INTO `area` (`descricao`, `ts`) VALUES
('zona sul', '2017-10-06 01:17:55'),
('zona leste', '2017-10-06 01:18:49'),
('tijuca', '2017-10-06 01:19:39'),
('andarai', '2017-10-06 01:20:15');

CREATE TABLE IF NOT EXISTS `cliente` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `nome` text(80) NOT NULL,
  `id_area` integer(6) NOT NULL,
  `cep` int(8) DEFAULT NULL,
  `endereco` text(80) NOT NULL,
  `ponto_de_referencia` text(40) DEFAULT NULL,
  `ddi` integer(3) NOT NULL DEFAULT '55',
  `ddd` integer(3) NOT NULL DEFAULT '21',
  `telefone` integer(9) DEFAULT NULL,
  `email` text(80) DEFAULT NULL,
  `flg_aviso_gas_final` integer(1) NOT NULL DEFAULT '1',
  `ts` text NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(id_area) REFERENCES area(id)
) ;

INSERT INTO `cliente` (`nome`, `id_area`, `cep`, `endereco`, `ponto_de_referencia`, `ddi`, `ddd`, `telefone`, `email`, `flg_aviso_gas_final`, `ts`) VALUES
('Jussara', 3, 99998, 'Rua Santo Agostinho, 50 - Casa 2', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:18'),
('Cosma', 4, 99998, 'Rua Andarai, 471', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:24'),
('Nei, pai do gugu', 4, 99998, 'Rua Andarai, 362 - Casa 9', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:27'),
('Alexandre', 4, 99998, 'Avenida Engenheiro Richard, 160 / 303', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:37'),
('Waldir', 4, 99998, 'Rua Anajatuba, 134', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:38'),
('Leila', 4, 99998, 'Rua Andarai, 563', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:41'),
('Jisiane', 3, 99998, 'Rua Monsenhor Batistoni, 355', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:41'),
('Valeria', 3, 99998, 'Rua Barão de Mesquita, 857', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:41'),
('Joana', 4, 99998, 'Rua Juparanã, 10', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:41'),
('Regina', 4, 99998, 'Rua Andarai, 471 - Casa 27', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:41'),
('Casa do cobra', 4, 99998, 'Rua Andarai, 513', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:42'),
('Denise ou Ronaldo', 4, 99998, 'Rua Ernesto de Souza, 133 treiler', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:42'),
('Elisa', 4, 99998, 'Rua Andarai, 294', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Dona Norma, Adriana irmã Slecia', 4, 99998, 'Rua Santo Estevão, 56', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Soraia', 4, 99998, 'Rua Leopoldo, 867', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Ana Paula / Paulo', 4, 99998, 'Rua Andarai, 471 - Casa 3', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Carlos', 4, 99998, 'Rua Itabaiana, 82', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Gil', 4, 99998, 'Rua Paula Brito, 706 A / 101', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Gideon', 4, 99998, 'Rua Leopoldo, 937 - Casa 2', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Andreia', 4, 99998, 'Rua Santo Agostinho, 51 - Casa 1', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Roberto', 4, 99998, 'Rua Anajatuba, 144 / 102', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Andreia', 4, 99998, 'Rua Santo Agostinho, 104 - Casa 12', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Romada / irmã do pia', 4, 99998, 'Rua Santo Agostinho, 532 - Casa 6', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Geovana', 4, 99998, 'Rua Juparanã, 10 / 401', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Ademar', 3, 99998, 'Rua Maria Amalia, 701', 'Uruguai', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Leandro', 3, 99998, 'Rua Uruguai, 455 F', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Adir', 4, 99998, 'Rua Santo Agostinho, 55', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Daniele', 4, 99998, 'Rua Santo Agostinho, 183', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Viviane', 4, 99998, 'Rua Santo Agostinho, 190 F', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Fabiano', 4, 99998, 'Rua Leopoldo, 937 - Casa 9', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Waldir', 4, 99998, 'Rua Ferreira Pontes, 104 - Casa 1', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:43'),
('Mica', 4, 99998, 'Rua Andarai, 471 - Casa 7', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Roseane', 4, 99998, 'Rua Santo Estevão, 87 / 404', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Filha Rosineide', 4, 99998, 'Rua Andarai, 513', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Celina', 4, 99998, 'Rua Leopoldo, 959 - Casa 35', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Marcos', 4, 99998, 'Rua Leopoldo, 990', 'Lado salão', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Francisca', 4, 99998, 'Rua Andarai, 401 - Casa 5', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Chiquinha', 4, 99998, 'Rua Leopoldo, 990', 'Frente Daniel', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Thiago', 4, 99998, 'Rua Andarai, 471 - Casa 26', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Alcina', 4, 99998, 'Rua Andarai, 676', 'Entrada pela Santo Agostinho', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Gesica', 4, 99998, 'Rua Andarai, 401 - Casa 6', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Luzinete', 4, 99998, 'Rua Santo Agostinho, 463 F', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Luiz', 4, 99998, 'Rua Leopoldo, 1507 - Casa 4', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Mauricio', 4, 99998, 'Rua Andarai, 419 - Casa 4', 'Entrada pela Santo Agostinho', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Carlos', 4, 99998, 'Rua Andarai, 513', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Diogo', 3, 99998, 'Travessa Matilde, 25 - Casa 7', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Simone', 4, 99998, 'Rua Engenheiro Morsing, 163', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Raquel', 4, 99998, 'Rua Andarai, 459 - Casa 202', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Joana', 4, 99998, 'Rua Santo Estevão, 87 / 301', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Gustavo', 4, 99998, 'Rua Santo Estevão, 128 - Casa 9', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Sabrina', 4, 99998, 'Rua Leopoldo, 110 - Casa 3', 'Antigo bar do frango assado', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Matheus', 4, 99998, 'Rua Maria Amalia, 713 - Casa 2', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Zé Carlos', 4, 99998, 'Rua Leopoldo,1121', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Edimar', 4, 99998, 'Rua Caçapava, 333', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Nilza', 4, 99998, 'Rua Santo Agostinho, 445 - Casa 104', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Nicolle', 4, 99998, 'Rua Leopoldo, 1165 - Casa 4', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Julieta', 4, 99998, 'Rua Ferreira Pontes, 764', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Leo', 4, 99998, 'Rua Santo Agostinho, 463 - Casa 102', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Nelson', 4, 99998, 'Rua Ferreira Pontes, 400 - Casa 4', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Andreia', 4, 99998, 'Rua Borda do Mato, 937 - Casa 1', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Rose', 4, 99998, 'Rua Santo Agostinho, 463 - Casa 103', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Beth', 4, 99998, 'Rua Engenheiro Morsing, 115', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Eliclene', 4, 99998, 'Rua Andarai, 70', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Evandro', 4, 99998, 'Rua Doná Florinda, 19', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Andre', 4, 99998, 'Rua Andarai, 471 - Casa 25', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Conceição', 4, 99998, 'Rua França Júnior, 476 - Casa 2', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Ladir', 4, 99998, 'Rua Santo Agostinho, 463 - Apto 201', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Cerli', 4, 99998, 'Rua Anajatuba, 106 - Casa 8', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Lucia', 4, 99998, 'Rua Paula Brito, 101 - Apto 104', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Célia', 4, 99998, 'Rua Borda do Mato, 420 - Casa 1', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Jurema', 4, 99998, 'Rua Gonzaga Bastos, 230 - Casa 12', 'Vila Isabel', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Alecia', 4, 99998, 'Rua Borda do Mato, 939 - Casa 4', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Pamela', 4, 99998, 'Rua Sousa Cruz, 67 - Apto 404', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Marilene', 4, 99998, 'Rua Santo Agostinho, 463', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Rose', 4, 99998, 'Rua Andarai, 513', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Maria', 4, 99998, 'Rua Andarai, 635 - Casa 11', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Sonia ou Cleide', 3, 99998, 'Rua Uruguai, 488 A - Casa 8', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Nilson', 4, 99998, 'Rua Andarai, 513', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Vanessa', 4, 99998, 'Rua Andarai, 19 - Casa 9', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Shirley', 4, 99998, 'Rua Ferreira Pontes, 769', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Jorge', 4, 99998, 'Rua Caçapava, 314 - Casa 3', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Tais', 4, 99998, 'Rua Campinas, 12 - Apto 101', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Celso', 4, 99998, 'Rua Santo Agostinho, 104 - Casa 4', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Adelson', 4, 99998, 'Rua Santo Estevão, 186 - Casa 1', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Sem nome', 3, 99998, 'Avenida Paulo de Frontim, 157 - 2o andar', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Rosangela', 4, 99998, 'Rua Andarai, 401 - Casa 2', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Leila', 4, 99998, 'Rua Andarai, 563 - Apto 302', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Mauricio', 4, 99998, 'Rua Ferreira Pontes, 753 - Casa 10', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:44'),
('Lucia', 4, 99998, 'Rua Andarai, 475', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:45'),
('José Mauricio', 4, 99998, 'Rua Leopoldo, 862 - Casa 15', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:45'),
('Luiz', 4, 99998, 'Rua Leopoldo, 1057 - Casa 4', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:45'),
('Wladimir', 4, 99998, 'Rua Ferreira Pontes, 104 - Casa 1', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:45'),
('Vilma', 4, 99998, 'Rua Conselheiro Paranaguá, 73 A - Casa 4', 'Vila Isabel', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:45'),
('Sebastião', 3, 99998, 'Rua Santa Luíza, 432', 'Maracanã', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:45'),
('Aurea', 4, 99998, 'Rua Pontes Corrêa, 260 - Apto 202', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:45'),
('Carequinha da kombi', 4, 99998, 'Rua Leopoldo, 1204', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:45'),
('Penha', 4, 99998, 'Rua Santo Agostinho, 183', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-12 22:49:45'),
('Beto', 4, 99997, 'Rua Andarai, 670 - Casa 4', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Janete', 4, 99997, 'Rua Andarai, 163 - Apto 106', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Maria José', 4, 99997, 'Rua Maria Amalia, 290 - Apto 201', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Bar do Juilo', 3, 99997, 'Rua Barão de Mesquita', 'Frente chaveiro', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Nicole', 4, 99997, 'Rua Leopoldo, 1145 - Casa 4', 'Frente Lixeira', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Maura', 4, 99997, 'Rua Leopoldo, 937', 'Frente Lixeira', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Clauida', 4, 99997, 'Rua Engenheiro Morsing, 80 - Apto 102', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Ari', 4, 99997, 'Rua Engenheiro Morsing, 80 - Apto 101', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Sindicato', 4, 99997, 'Rua Uberaba, 36', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Deise', 4, 99997, 'Rua Paula Brito, 711 - C A', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Padaria Maria Amalia', 4, 99997, 'Padaria Maria Amalia', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Vanusa', 4, 99997, 'Rua Andarai, 119 - Casa 9', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Elideane', 4, 99997, 'Rua Andarai, 670 - Casa 1', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Andarai', 4, 99997, 'Rua Andarai, 459 - Apto 102', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Solange', 4, 99997, 'Rua Leopoldo, 1115 - Casa 2', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Cerlison', 4, 99997, 'Rua Andarai, 163 - Apto 103', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Maria José', 4, 99997, 'Rua Andarai, 151', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Fani', 4, 99997, 'Rua Raja Gabaglia, 64', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Lucia', 4, 99997, 'Rua Andarai, 375', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Rose', 4, 99997, 'Rua Andarai, 513', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Jordana', 4, 99997, 'Rua Leopoldo, 862', 'Frente Lixeira', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:21'),
('Fernanda', 3, 99997, 'Rua Caetano de Campos, 94 - Casa 102', 'Alto da Boa Vista', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Tartaruga', 4, 99997, 'Rua Andarai, 671', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Fernanda', 4, 99997, 'Rua Alfredo Pujal, 64 Fundos', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Xavier', 4, 99997, 'Rua Leopoldo, 157 - Casa 2', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Lidiana', 4, 99997, 'Rua Leopoldo, 1209', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Alexandra, Bar do gordo', 4, 99997, 'Rua Leopoldo', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Maria', 3, 99997, 'Rua Santa Luíza, 432 Fundos', 'Maracanã', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Lusinete', 4, 99997, 'Rua Santo Agostinho, 463', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Andre', 4, 99997, 'Rua Andarai, 471 - Casa 25', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Erica', 4, 99997, 'Rua Andarai, 459 - Apto 303', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Jorge', 4, 99997, 'Rua Caçapava, 314 - Casa 3', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Ana', 4, 99997, 'Rua Andarai, 95 - Casa 7 Fundos', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Jacare', 4, 99997, 'Rua Leopoldo', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Maura', 4, 99997, 'Rua Leopoldo, 937', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Carine', 4, 99997, 'Rua Adolfo Caminha, 52 Fundos', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Neusa', 4, 99997, 'Rua Andarai, 401', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Wladimir', 4, 99997, 'Rua Ferreira Pontes, 104 - Casa 1', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Carina - Mario Preto', 4, 99997, 'Rua Andarai, 401 Fundos', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Joaquim', 4, 99997, 'Rua Andarai, 348', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Mariane', 4, 99997, 'Rua Leopoldo, 937 - Casa 9', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Daniele', 4, 99997, 'Rua Caçapava, 306 - Casa 1', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Margarida', 4, 99997, 'Rua Divinéia Peterson, 152', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Luiz', 4, 99997, 'Rua Maria Amalia, 265', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Antonia', 3, 99997, 'Rua Santo Agostinho, 387 - Casa 1', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Valeria', 3, 99997, 'Rua Barão de Mesquita, 857- Grupo 1 - Apto 102', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('João', 3, 99997, 'Rua Santo Agostinho, 104 - Casa 12', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Maria Luiza', 4, 99997, 'Rua Andarai, 401', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Amanda', 3, 99997, 'Rua Carvalho Alvim, 529', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Conceição', 4, 99997, 'Rua Conselheiro Paranaguá, 73 A', 'Vila Isabel', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Geovana', 4, 99997, 'Rua Juparanã, 10 / 401', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Murilo', 4, 99997, 'Rua Leopoldo, 937 - Casa 10', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Flaclemi', 4, 99997, 'Rua Ferreira Pontes, 101 - Casa 1', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Rosangela', 3, 99997, 'Rua Barão de Mesquita, 743 - Apto 705', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Isabel', 4, 99997, 'Rua Conselheiro Paranaguá, 73 A - Casa 8', 'Vila Isabel', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Isac', 4, 99997, 'Rua Maria Amalia, 681', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Barbara', 4, 99997, 'Rua Paula Brito, 476 - Apto 201', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Patricia', 4, 99997, 'Rua Andarai, 513', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Waldir', 4, 99997, 'Rua Anajatuba, 134', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Waldecir', 4, 99997, 'Rua Anajatuba, 134 - Casa 1', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Marilene', 4, 99997, 'Rua Ernesto de Souza, 133 Bloco 1 - Casa 27', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Beto', 4, 99997, 'Rua Leopoldo, 1165 - Casa 3', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Sonia', 4, 99997, 'Rua Marechal Jofre, 73 - Casa 5', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Fernanda', 3, 99997, 'Rua Santo Agostinho, 445 - Casa 4', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Restaurante', 3, 99997, 'Rua Bom Pastor, 357', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Gil Material', 4, 99997, 'Rua Borda do Mato, 427 F', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('José', 4, 99997, 'Rua Leopoldo, 990', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Clicio', 4, 99997, 'Rua Conselheiro Paranaguá, 73 A - Casa 10', 'Vila Isabel', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Penha', 4, 99997, 'Rua Caçapava, 301 - Casa 4', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Lucia', 4, 99997, 'Rua Leopoldo, 1077 - Casa 8', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Nanci', 4, 99997, 'Rua Andarai, 459 - Apto 102', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Luis Claudio', 4, 99997, 'Rua Pontes Correia, 260 Fundos', 'Rua Maxwell, paralela Rua Uruguai', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Soraia', 4, 99997, 'Rua Leopoldo, 867', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Iara', 4, 99997, 'Rua Santo Estevão, 134', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Bar do Cabeludo', 4, 99997, 'Rua Borda do Mato, 410', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Jose', 3, 99997, 'Rua Santo Agostinho, 335', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Gabriele', 3, 99997, 'Rua Bom Pastor, 89 - Casa 11 - 201', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Luciano', 4, 99997, 'Rua Ferreira Pontes, 688', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Claudio', 4, 99997, 'Rua Anajatuba, 133 - Casa 1', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Fatima', 4, 99997, 'Rua Anajatuba, 134 - Casa 2', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Geronimo', 4, 99997, 'Rua Leopoldo, 937', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Elidiana', 4, 99997, 'Rua Andarai, 470 - Casa 1', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Mauro', 3, 99997, 'Rua Santo Agostinho, 14', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Solange', 4, 99997, 'Rua Leopoldo, 1155 - Casa 2', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Ana', 4, 99997, 'Praça Nobel, 148 - Apto 104', 'Grajaú', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Rosa', 4, 99997, 'Rua Teodoro da Silva, 363 - Apto 201', 'Vila Isabel', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Anderson', 4, 99997, 'Rua Doná Florinda, 32 A', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Ana Rita', 3, 99997, 'Rua Clemente Falcão, 153 - Apto 102', 'Extra / Maracanã', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Ana Rita', 3, 99997, 'Rua Dona Angélica, 152', 'Rua Mariz e Barros', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Flavia', 4, 99997, 'Rua Andarai, 660', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Viviane', 3, 99997, 'Rua Carlos Vasconcelos, 139 - Apto 301', 'Rua Mariz e Barros', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Aninha', 3, 99997, 'Rua Santo Agostinho, 530 - Casa 4', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Luis', 4, 99997, 'Rua Leopoldo, 57 - Casa 4', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Leila', 4, 99997, 'Rua Andarai, 563 - Apto 302', 'Hospital Federal Andarai ', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22'),
('Nilsa', 3, 99997, 'Rua Santo Agostinho, 445 - Casa 4', '', 55, 21, NULL, 'xxx@gmail.com', 1, '2017-11-23 13:56:22');

CREATE TABLE IF NOT EXISTS `cobertura` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `id_usuario` integer(4) NOT NULL,
  `id_area` integer(6) NOT NULL,
  `ts` text NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(id_usuario) REFERENCES usuario(id),
  FOREIGN KEY(id_area) REFERENCES area(id)
) ;

INSERT INTO `cobertura` (`id_usuario`, `id_area`, `ts`) VALUES
(3, 2, '2017-10-06 01:20:50'),
(4, 1, '2017-10-06 01:21:01'),
(6, 4, '2017-10-06 01:21:11'),
(5, 3, '2017-10-06 01:21:21');

CREATE TABLE IF NOT EXISTS `glp` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `descricao` text(20) NOT NULL,
  `glp_default` integer(1) NOT NULL,
  `ts` text NOT NULL DEFAULT CURRENT_TIMESTAMP
) ;

INSERT INTO `glp` (`descricao`, `glp_default`, `ts`) VALUES
('P13', 1, '2017-10-06 01:27:56'),
('P45', 0, '2017-10-06 01:27:56'),
('P13 completo', 0, '2017-10-06 01:28:32');

CREATE TABLE IF NOT EXISTS `glp_preco_area` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `id_glp` integer(4) NOT NULL,
  `id_area` integer(6) NOT NULL,
  `preco` decimal(10,2) NOT NULL,
  `ts` text NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(id_area) REFERENCES area(id),
  FOREIGN KEY(id_glp) REFERENCES glp(id)
) ;

INSERT INTO `glp_preco_area` (`id_glp`, `id_area`, `preco`, `ts`) VALUES
(1, 4, '72.00', '2017-10-06 01:29:36'),
(1, 3, '72.00', '2017-10-06 01:29:50'),
(1, 2, '77.00', '2017-10-06 01:30:04'),
(1, 1, '77.00', '2017-10-06 01:30:19'),
(2, 4, '250.00', '2017-10-06 01:30:53'),
(2, 3, '250.00', '2017-10-06 01:31:08'),
(2, 2, '250.00', '2017-10-06 01:31:29'),
(2, 1, '250.00', '2017-10-06 01:31:41'),
(3, 4, '220.00', '2017-10-06 01:31:58'),
(3, 3, '220.00', '2017-10-06 01:32:11'),
(3, 2, '220.00', '2017-10-06 01:32:24'),
(3, 1, '220.00', '2017-10-06 01:32:38');

CREATE TABLE IF NOT EXISTS `pedido` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `id_cliente` integer(11) NOT NULL,
  `quantidade` integer(4) NOT NULL,
  `id_glp_preco_area` integer(11) NOT NULL,
  `valor_pedido` decimal(10,2) NOT NULL,
  `data_pedido` date NOT NULL,
  `flg_prioridade_pedido` integer(1) NOT NULL,
  `flg_origem_pedido` integer(1) NOT NULL,
  `flg_hora` integer(2) NOT NULL DEFAULT '0',
  `flg_status` integer(1) NOT NULL DEFAULT '0',
  `data_status` datetime DEFAULT NULL,
  `ts` text NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(id_cliente) REFERENCES cliente(id),
  FOREIGN KEY(id_glp_preco_area) REFERENCES glp_preco_area(id)
) ;

CREATE TABLE IF NOT EXISTS `usuario` (
  `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  `login` text(20) NOT NULL,
  `senha` text(20) NOT NULL,
  `nome` text(40) NOT NULL,
  `flg_tipo_usuario` integer(4) NOT NULL,
  `ts` text NOT NULL DEFAULT CURRENT_TIMESTAMP
) ;

INSERT INTO `usuario` (`login`, `senha`, `nome`, `flg_tipo_usuario`, `ts`) VALUES
('evandro', 'teste123', 'Evandro', 0, '2018-01-11 11:50:39'), --usuario Adm
('eduardo', 'teste123', 'Eduardo', 1, '2018-01-11 11:50:39'), --usuario light Deposito/Motorista
('bidu', 'teste123', 'Bidu', 2, '2018-01-11 11:50:39'), --usuario Deposito
('lima', 'teste123', 'Lima', 2, '2018-01-11 11:50:45'), --usuario Deposito
('bruno', 'teste123', 'Bruno', 3, '2017-10-06 01:15:24'), --usuario Motorista
('wilian', 'teste123', 'Wilian', 3, '2017-10-06 01:15:52'), --usuario Motorista
('paulo', 'teste123', 'Paulo', 3, '2017-10-06 01:16:34'), --usuario Motorista
('victor', 'teste123', 'Victor', 3, '2017-10-06 01:17:13'); --usuario Motorista

CREATE INDEX IF NOT EXISTS idx_area_key_descricao ON area(descricao);

CREATE INDEX IF NOT EXISTS idx_cliente_key_id_area ON cliente(id_area);
CREATE INDEX IF NOT EXISTS idx_cliente_key_endereco ON cliente(endereco);
CREATE INDEX IF NOT EXISTS idx_cliente_key_nome ON cliente(nome);
CREATE INDEX IF NOT EXISTS idx_cliente_key_telefone ON cliente(telefone);

CREATE INDEX IF NOT EXISTS idx_cobertura_key_id_usuario_vs_id_area ON cobertura(id_usuario, id_area);
CREATE INDEX IF NOT EXISTS idx_cobertura_key_id_area_vs_id_usuario ON cobertura(id_area, id_usuario);

CREATE INDEX IF NOT EXISTS idx_glp_key_descricao ON glp(descricao);

CREATE INDEX IF NOT EXISTS idx_glp_preco_area_key_id_glp_vs_id_area ON glp_preco_area(id_glp, id_area);

CREATE INDEX IF NOT EXISTS idx_pedido_key_id_cliente ON pedido(id_cliente);
CREATE INDEX IF NOT EXISTS idx_pedido_key_data_pedido ON pedido(data_pedido);
CREATE INDEX IF NOT EXISTS idx_pedido_key_ts ON pedido(ts);
CREATE INDEX IF NOT EXISTS idx_pedido_key_id_glp_preco_area ON pedido(id_glp_preco_area);
CREATE INDEX IF NOT EXISTS idx_pedido_key_flg_origem_pedido ON pedido(flg_origem_pedido);

CREATE INDEX IF NOT EXISTS idx_usuario_key_login ON usuario(login);


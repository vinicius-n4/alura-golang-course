create table produtos(
    id serial primary key,
    nome varchar,
    descricao varchar,
    preco float,
    quantidade int
);

INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES
('Laptop', 'i7, 512gb SSD, 16gb RAM', 1070.50, 3)

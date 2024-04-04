CREATE TABLE equations(
    id SERIAL PRIMARY KEY,
    equation TEXT,
    answer NUMERIC,
    difficulty INTEGER
);

INSERT INTO equations (equation, answer, difficulty) VALUES
    ('2 + x = 5', 3, 1),
    ('3x = 9', 3, 1),
    ('7 - y = 3', 4, 1),
    ('x + 4 = 9', 5, 1),
    ('2x + 3 = 11', 4, 1),

    ('2x + y = 10', 3, 2),
    ('4x - 3y = 5', 4, 2),
    ('3x + 2y = 16', 4, 2),
    ('2x - y = 8', 4, 2),
    ('x + y = 7', 4, 2),

    ('3x - 2y = 4', 3, 3),
    ('2x^2 + 3y = 11', 2, 3),
    ('4x + 2y = 14', 3, 3),
    ('x^2 - y = 3', 2, 3),
    ('3x + 4y = 22', 4, 3);


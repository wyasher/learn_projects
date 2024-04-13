CREATE TABLE tod.todo_list(
                              id SERIAL PRIMARY KEY,
                              title VARCHAR(150) NOT NULL
);

CREATE TABLE tod.todo_item(
                              id SERIAL PRIMARY KEY,
                              title VARCHAR(150) NOT NULL,
                              checked BOOLEAN NOT NULL DEFAULT FALSE,
                              list_id INTEGER NOT NULL,
                              FOREIGN KEY(list_id) REFERENCES tod.todo_list(id)
);

INSERT INTO tod.todo_list (title) VALUES ('清单1'), ('清单2');

INSERT INTO tod.todo_item (title, list_id) VALUES
                                               ('清单项目 1', 1),
                                               ('清单项目 2', 1),
                                               ('清单项目 A', 2);
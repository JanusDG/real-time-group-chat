CREATE TABLE _user (
  user_id varchar(100) UNIQUE NOT NULL,
  username varchar(100) UNIQUE NOT NULL,
  name varchar(100) NOT NULL,
  surname varchar(100) NOT NULL,
  password varchar(100) NOT NULL,
  PRIMARY KEY (user_id)
);

CREATE TABLE _group (
  group_id varchar(100) UNIQUE NOT NULL,
  group_name varchar(100) NOT NULL,
  PRIMARY KEY (group_id)
);

CREATE TABLE _usergroup (
  group_id varchar(100) NOT NULL,
  user_id varchar(100) NOT NULL,
  PRIMARY KEY (group_id, user_id),
  FOREIGN KEY (group_id) REFERENCES _Group(group_id),
  FOREIGN KEY (user_id) REFERENCES _User(user_id)
);

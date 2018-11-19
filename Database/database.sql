set sql_mode='';

-- DROP DATABASE IF EXISTS mntbudget;

CREATE DATABASE IF NOT EXISTS mntbudget;

USE mntbudget;

CREATE TABLE IF NOT EXISTS profile (
  profile_id      INT UNSIGNED NOT NULL AUTO_INCREMENT,
  username        VARCHAR(255) NOT NULL,
  preferences     TEXT CHARACTER SET utf8mb4,

  PRIMARY KEY ( profile_id )
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

CREATE TABLE IF NOT EXISTS salary (
  profile_id      INT UNSIGNED NOT NULL,
  description     VARCHAR(255) NOT NULL,
  bruto_month     DECIMAL(13,2),
  bruto_year      DECIMAL(13,2),
  estimated_net   DECIMAL(13,2),
  monthly_tax     DECIMAL(13,2),
  yearly_tax      DECIMAL(13,2),
  dutch_30        TINYINT DEFAULT 0,

  PRIMARY KEY (profile_id,description)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;


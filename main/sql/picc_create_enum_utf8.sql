
/*==============================================================*/
/*                       create enums                           */
/*==============================================================*/

CREATE TYPE relation_type AS ENUM('本人','村民与村委');

CREATE TYPE insurance_type AS ENUM('个人投保', '团体投保');

CREATE TYPE loss_level_type AS ENUM('轻微','一般','较重','严重');

CREATE TYPE insurance_status_type AS ENUM('正常','审核中','已审核','出险中','出险完毕','过期');


INSERT INTO `organizations` (`id`, `name`, `code`, `mission`, `vision`, `value`) VALUES
('01GZR0ZEMY9W4GP74E3G6GYEWY', '新卒研修部', 'NewGraduateTraining', 'ミッション', 'ビジョン', 'バリュー'),
('01GZR2TYVGFJKWH35BF2J5Z38E', '新卒研修部_後編', 'NewGraduateTraining2', 'ミッション', 'ビジョン', 'バリュー');

INSERT INTO `users` (`id`, `cognito_id`, `name`, `email`) VALUES ('01GZT0HJAX3CM9P1V66T9GYM95', '28f155b7-860f-49c4-b059-5eed51a4ce4c', 'test.local', 'test.local@gmail.com');
INSERT INTO `r_organization_users` (`user_id`, `organization_id`, `role`) VALUES ('01GZT0HJAX3CM9P1V66T9GYM95', '01GZR2TYVGFJKWH35BF2J5Z38E', 1);

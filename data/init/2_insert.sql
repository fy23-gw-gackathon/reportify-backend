INSERT INTO `organizations` (`id`, `name`, `code`, `mission`, `vision`, `value`) VALUES
('01GZR2TYVGFJKWH35BF2J5Z38E', '技術統括部エンジニアリング室新卒研修', 'fy23-eng-training', '一人ひとりに想像を超えるDelightを', 'DeNAは、インターネットやAIを自在に駆使しながら 一人ひとりの人生を豊かにするエンターテインメント領域と 日々の生活を営む空間と時間をより快適にする社会課題領域の 両軸の事業を展開するユニークな特性を生かし 挑戦心豊かな社員それぞれの個性を余すことなく発揮することで 世界に通用する新しいDelightを提供し続けます', 'DeNA Promise');

INSERT INTO `users` (`id`, `cognito_id`, `name`, `email`) VALUES ('01GZT0HJAX3CM9P1V66T9GYM95', '28f155b7-860f-49c4-b059-5eed51a4ce4c', 'test.local', 'test.local@gmail.com');
INSERT INTO `r_organization_users` (`user_id`, `organization_id`, `role`) VALUES ('01GZT0HJAX3CM9P1V66T9GYM95', '01GZR2TYVGFJKWH35BF2J5Z38E', 1);

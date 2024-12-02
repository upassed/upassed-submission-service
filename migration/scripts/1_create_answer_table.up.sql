create table if not exists answer (
    id uuid primary key,
    student_username varchar not null,
    form_id uuid not null,
    question_id uuid not null,
    answer_id uuid not null
);

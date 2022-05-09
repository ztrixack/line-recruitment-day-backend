CREATE TABLE public.elections (id BIGINT, enable BOOLEAN, state TEXT);

CREATE TABLE public.candidates (
	id BIGINT,
	name TEXT,
	dob TEXT,
	bio_link TEXT,
	image_link TEXT,
	policy TEXT
);

CREATE TABLE public.candidate_votes (candidate_id BIGINT, voted_count BIGINT);

-- TRUNCATE TABLE public.elections;

INSERT INTO
	public.elections(id, enable, state)
VALUES
	(1, false, 'solicit');

-- TRUNCATE TABLE public.candidates;

INSERT INTO
	public.candidates(id, name, dob, bio_link, image_link, policy)
VALUES
	(
		1,
		'Elon Mask',
		'June 28, 1971',
		'<https://en.wikipedia.org/wiki/Elon_Musk>',
		'<http://placekitten.com/480/640>',
		'Lorem Ipsum is simply dummy text of the printing and typesetting industry.'
	);

INSERT INTO
	public.candidates(id, name, dob, bio_link, image_link, policy)
VALUES
	(
		2,
		'Brown',
		'August 8, 2011',
		'<https://line.fandom.com/wiki/Brown>',
		'<http://placekitten.com/480/640>',
		'Lorem Ipsum is simply dummy text of the printing and typesetting industry.'
	);

INSERT INTO
	public.candidates(id, name, dob, bio_link, image_link, policy)
VALUES
	(
		3,
		'LINE Brown',
		'August 8, 2011',
		'<https://line.fandom.com/wiki/Brown>',
		'<http://placekitten.com/480/640>',
		'Lorem Ipsum is simply dummy text of the printing and typesetting industry.'
	);

INSERT INTO
	public.candidates(id, name, dob, bio_link, image_link, policy)
VALUES
	(
		4,
		'Jeff Bezos',
		'January 12, 1964',
		'<https://en.wikipedia.org/wiki/Jeff_Bezos>',
		'<http://placekitten.com/480/640>',
		'Lorem Ipsum is simply dummy text of the printing and typesetting industry.'
	);

INSERT INTO
	public.candidates(id, name, dob, bio_link, image_link, policy)
VALUES
	(
		5,
		'Bill Gates',
		'October 28, 1955',
		'<https://en.wikipedia.org/wiki/Bill_Gates>',
		'<http://placekitten.com/480/640>',
		'Lorem Ipsum is simply dummy text of the printing and typesetting industry.'
	);

INSERT INTO
	public.candidates(id, name, dob, bio_link, image_link, policy)
VALUES
	(
		6,
		'Mark Zuckerberg',
		'May 14, 1984',
		'<https://en.wikipedia.org/wiki/Mark_Zuckerberg>',
		'<http://placekitten.com/480/640>',
		'Lorem Ipsum is simply dummy text of the printing and typesetting industry.'
	);

INSERT INTO
	public.candidates(id, name, dob, bio_link, image_link, policy)
VALUES
	(
		7,
		'Larry Page',
		'March 26, 1973',
		'<https://en.wikipedia.org/wiki/Larry_Page>',
		'<http://placekitten.com/480/640>',
		'Lorem Ipsum is simply dummy text of the printing and typesetting industry.'
	);

INSERT INTO
	public.candidates(id, name, dob, bio_link, image_link, policy)
VALUES
	(
		8,
		'Sergey Brin',
		'August 21, 1973',
		'<https://en.wikipedia.org/wiki/Sergey_Brin>',
		'<http://placekitten.com/480/640>',
		'Lorem Ipsum is simply dummy text of the printing and typesetting industry.'
	);

-- TRUNCATE TABLE public.candidate_votes;

INSERT INTO
	public.candidate_votes(candidate_id, voted_count)
VALUES
	(1, 0);

INSERT INTO
	public.candidate_votes(candidate_id, voted_count)
VALUES
	(2, 0);

INSERT INTO
	public.candidate_votes(candidate_id, voted_count)
VALUES
	(3, 0);

INSERT INTO
	public.candidate_votes(candidate_id, voted_count)
VALUES
	(4, 0);

INSERT INTO
	public.candidate_votes(candidate_id, voted_count)
VALUES
	(5, 0);

INSERT INTO
	public.candidate_votes(candidate_id, voted_count)
VALUES
	(6, 0);

INSERT INTO
	public.candidate_votes(candidate_id, voted_count)
VALUES
	(7, 0);

INSERT INTO
	public.candidate_votes(candidate_id, voted_count)
VALUES
	(8, 0);
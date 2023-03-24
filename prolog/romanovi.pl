parent("Михаил", "Алексей").
parent("Евдокия", "Алексей").
parent("Алексей", "Иоанн").
parent("Мария", "Иоанн").
parent("Алексей", "Софья").
parent("Мария", "Софья").
parent("Алексей", "Феодор").
parent("Мария", "Феодор").
parent("Алексей", "Петр").
parent("Наталья", "Петр").
parent("Петр", "Анна").
parent("Екатерина", "Анна").
parent("Петр", "Елизавета").
parent("Екатерина", "Елизавета").

man("Михаил").
man("Алексей").
man("Иоанн").
man("Феодор").
man("Петр").

woman("Евдокия").
woman("Мария").
woman("Софья").
woman("Наталья").
woman("Екатерина").
woman("Анна").
woman("Елизавета").

mother(Who, Whom) :-
    parent(Who, Whom),
    woman(Who).
father(Who, Whom) :-
    parent(Who, Whom),
    man(Who).
son(Who, Whom) :-
    parent(Whom, Who),
    man(Who).
daughter(Who, Whom) :-
    parent(Whom, Who),
    woman(Who).

grandparent(Who, Whom) :-
    parent(Who, Par),
    parent(Par, Whom).
grandfather(Who, Whom) :-
    parent(Who, Par),
    parent(Par, Whom),
    man(Who).
grandmother(Who, Whom) :-
    parent(Who, Par),
    parent(Par, Whom),
    woman(Who).
brother(Who, Whom) :-
    (   parent(P1, Whom),
        son(Who, P1),
        Whom\==Who,
        \+ father(_, Who)
    ;   parent(P1, Whom),
        son(Who, P1),
        Whom\==Who,
        \+ mother(_, Who)
    ;   father(P2, Who),
        father(P3, Whom),
        P2\==P3,
        parent(P1, Whom),
        son(Who, P1),
        Whom\==Who
    ;   mother(P2, Who),
        mother(P3, Whom),
        P2\==P3,
        parent(P1, Whom),
        son(Who, P1),
        Whom\==Who
    ;   father(P2, Who),
        father(P2, Whom),
        mother(P1, Who),
        mother(P1, Whom),
        parent(P1, Whom),
        son(Who, P1),
        Whom\==Who
    ).
sister(Who, Whom) :-
    (   parent(P1, Whom),
        daughter(Who, P1),
        Whom\==Who,
        \+ father(_, Who)
    ;   parent(P1, Whom),
        daughter(Who, P1),
        Whom\==Who,
        \+ mother(_, Who)
    ;   father(P2, Who),
        father(P3, Whom),
        P2\==P3,
        parent(P1, Whom),
        daughter(Who, P1),
        Whom\==Who
    ;   mother(P2, Who),
        mother(P3, Whom),
        P2\==P3,
        parent(P1, Whom),
        daughter(Who, P1),
        Whom\==Who
    ;   father(P2, Who),
        father(P2, Whom),
        mother(P1, Who),
        mother(P1, Whom),
        parent(P1, Whom),
        daughter(Who, P1),
        Whom\==Who
    ).
nephew(Who, Whom) :-
    brother(Br, Whom),
    son(Who, Br).
niece(Who, Whom) :-
    brother(Br, Whom),
    daughter(Who, Br).

grandson(Who, Whom) :-
    grandparent(Whom, Who),
    man(Who).
granddaughter(Who, Whom) :-
    grandparent(Whom, Who),
    woman(Who).

parents(P1, P2) :-
    parent(P1, Who),
    parent(P2, Who).

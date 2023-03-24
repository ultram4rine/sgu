genl(N2, N2, []) :- !.
genl(N1, N2, [N1|L]) :-
    N1>N2,
    N is N1-1,
    genl(N, N2, L).

insert(E, [], [E]).
insert(X, [Y|Rest], [X, Y|Rest]) :-
    X@>Y, !.
insert(X, [Y|L1], [Y|L2]) :-
    insert(X, L1, L2).

insert_all_pos_h(A, [], E) :-
    flatten([A|E], L1),
    write(L1).
insert_all_pos_h(A, [H|L], E) :-
    flatten([A, E, H|L], L1),
    write(L1),
    nl,
    insert_all_pos_h([A|H], L, E).
insert_all_pos([H|L], E) :-
    write([E, H|L]),
    nl,
    insert_all_pos_h(H, L, E).

sum([], 0).
sum([H|L], S) :-
    sum(L, S1),
    S is S1+H.

len([], 0).
len([_|L], X) :-
    len(L, X1),
    X is X1+1.

split(L, L1, L2) :-
    append(L1, L2, L),
    len(L1, N),
    len(L2, N).
split(L, L1, L2) :-
    append(L1, L2, L),
    len(L1, N),
    N1 is N+1,
    len(L2, N1).
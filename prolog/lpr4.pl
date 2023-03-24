pow(_, 0, 1).
pow(X, N, A) :-
    N1 is N-1,
    pow(X, N1, A1),
    A is A1*X.

fib(0, 0).
fib(1, 1).
fib(N, F) :-
    N1 is N-1,
    N2 is N-2,
    fib(N1, F1),
    fib(N2, F2),
    F is F1+F2.

kern(A, _, C, 1) :-
    write((A->C)),
    nl, !.
kern(A, B, C, N) :-
    N1 is N-1,
    kern(A, C, B, N1),
    write((A->C)),
    nl,
    kern(B, A, C, N1).

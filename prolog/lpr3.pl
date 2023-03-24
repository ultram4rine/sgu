s(X, Y) :-
    Res is 2*(X*X+Y*Y)/(X+Y),
    write("S="),
    write(Res).

is_even(X) :-
    (   write(X),
        write(" is "),
        M is X mod 2,
        M==0
    ->  write("even")
    ;   write("odd")
    ).

min_3(A, B, C) :-
    write("min is "),
    D is min(A, B),
    E is min(D, C),
    write(E).

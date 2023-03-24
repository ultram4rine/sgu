dict("Я", "I").
dict("ты", "you").
dict("спасибо", "thanks").

show_all :-
    dict(R, E),
    write(R),
    write(" "),
    write(E),
    write(" ").

r_e(W) :-
    dict(W, E),
    write(E).

e_r(W) :-
    dict(R, W),
    write(R).

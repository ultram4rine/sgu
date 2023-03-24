country("Россия", "Москва", 146, "Европа").
country("Польша", "Краков", 37, "Европа").
country("Чехия", "Прага", 10, "Европа").
country("Казахстан", "Астана", 19, "Азия").
country("Монголия", "Улан-Батор", 3, "Азия").

show_all :-
    country(C, DC, N, G),
    write(C),
    write(" "),
    write(DC),
    write(" "),
    write(N),
    write(" "),
    write(G).

n_g(M) :-
    country(C, DC, N, G),
    N>M,
    write(C),
    write(" "),
    write(DC),
    write(" "),
    write(N),
    write(" "),
    write(G),
    write(" ").

e_l_t(M) :-
    country(C, DC, N, "Европа"),
    N=<M,
    write(C),
    write(" "),
    write(DC),
    write(" "),
    write(N),
    write(" ").

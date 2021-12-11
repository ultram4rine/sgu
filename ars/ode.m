function task3
Tmax=25;
Tspan=(0:200)*(Tmax/200);

tic;
[T,Y]=ode113(@ODE_Fun, Tspan, [0; 0; 0; 0]);
toc;
figure; plot(T, Y(:,1));

tic;
[T,Y]=ode15s(@ODE_Fun, Tspan, [0; 0; 0; 0]);
toc;
figure; plot(T, Y(:,1));
end

function f=Fun(y)
f=[(0.02*y(2)*y(2))/(1+2*y(1)*y(1)+3*y(2)*y(2));
   (0.02*y(1)*y(1))/(1+3*y(1)*y(1)+y(1)*y(1));
   (2*y(2)*y(4))/(1+y(1)*y(1)+y(2)*y(2));
   -3*y(1)*y(3)];
end

function x=x_Fun(t)% Функция x(t) в правой части (1)
    x=[exp(-2*t)*cos(t); 1+exp(-t)*sin(t); 0; 0];
end

function F=ODE_Fun(t,y)
    A=[0.01    1    -0.01    0.01;
         -1 -0.5     0.01   -0.01;
        300 -230      700   1.0e5;
        120  350   -1.0e5  -1.0e5];
    F=A*y+x_Fun(t)+Fun(y);
end

function task4
R=10;
a=0.5;
xmesh=0.05*(0:20);
tspan=0.1*(0:10);
tspan(12:21)=1+0.2*(1:10);
tspan(22:25)=[3.5, 4, 4.5, 5];

sol = pdepe(0, @pdefun, @icfun, @bcfun, xmesh, tspan);

for k=2:size(tspan,2)
    fprintf('t=%f\n', tspan(k));
    UU=FFi(a+xmesh,tspan(k));
    disp('приближенное и точное решение');
    for j=1:size(xmesh,2)
        fprintf('%f\t%f\n', sol(k,j), UU(j));
    end
    disp('-------------------------------------------');
end
%---------------------------------------------------
%графическое отображение решения
figure; surf(xmesh,tspan,sol);
xlabel('x');
ylabel('t');

%---------------------------------
%Собственно уравнение в частных производных
     function [c,f,s]=pdefun(x,t,u,dudx)
     c=1;
     f=dudx/R;
     s=-( ( (2+cos(t))/2 ) * u * dudx ) + ( ( sin(t)/(2+cos(t)) ) * u );
     end
%-------------------------------
%Начальные условия
    function u = icfun(x)
    u=0;
    end
%-------------------------------
%Граничные условия
     function [pl,ql,pr,qr] = bcfun(xl,ul,xr,ur,t)
     pl=ul - ( (2*FFi(a,t)) / (2+cos(t)) );
     ql=0;
     pr=ur - ( (2*FFi(1+a,t)) / (2+cos(t)) );
     qr=0;
     end

%-----------------------------------------------------
    function u=FFi(x,t)
        if t>0.01
           T1=erfc(0.5*(x-t)*sqrt(R/t));
           T2=erfc(0.5*x*sqrt(R/t));
           u=T1./(T1+exp(0.25*R*(2*x-t)).*(1-T2));
        else
            u=0*x;
        end
    end
%--------------------------------
end
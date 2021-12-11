function task2
dt=0.01;
W=2*pi/dt;
N=10000;
dw=W/(2*N);
t=(0:2*N)*dt;

w1=37; w2=150;
A1=1; A2=0.7;

x=A1*f1(w1*t)+A2*f2(w2*t);
y=x+(A1+A2)*randn(size(t));
w=dw*(0:N-1);

X=fft(x);
figure; plot(w,abs(X(1:N)));

Y=fft(y);
figure; plot(w,abs(Y(1:N)));
end

function x=f1(t)
   x=sin((pi/2) * sin(t));
end

function x=f2(t)
   x=cos(t);
end
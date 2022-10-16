document.addEventListener("DOMContentLoaded", function (_event) {
  const portfolioSlider = new Swiper(".portfolio__slider", {
    // Default parameters
    slidesPerView: 1,
    spaceBetween: 20,
    loop: true,
    // Responsive breakpoints
    breakpoints: {
      // when window width is >= 320px
      320: {
        slidesPerView: 1,
        spaceBetween: 20,
      },
      // when window width is >= 480px
      480: {
        slidesPerView: 2,
        spaceBetween: 20,
      },
      // when window width is >= 640px
      640: {
        slidesPerView: 3,
        spaceBetween: 20,
      },
    },
  });
});

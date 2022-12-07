document.addEventListener("DOMContentLoaded", function (_event) {
  const portfolioSlider = new Swiper(".reviews__row", {
    // Default parameters
    slidesPerView: 1,
    spaceBetween: 30,
    loop: true,
    // Responsive breakpoints
    breakpoints: {
      // when window width is >= 1360px
      1360: {
        slidesPerView: 3,
        spaceBetween: 30,
      },
    },

    navigation: {
      nextEl: ".arrow-right",
      prevEl: ".arrow-left",
    },
  });
});

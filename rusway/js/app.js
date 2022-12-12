document.addEventListener("DOMContentLoaded", function (_event) {
  const portfolioSlider = new Swiper(".reviews__row", {
    // Default parameters
    slidesPerView: 1,
    spaceBetween: 30,
    loop: true,
    // Responsive breakpoints
    breakpoints: {
      // when window width is >= 1360px
      1400: {
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

function swapValues() {
  const tmp = document.getElementById("from").value;
  document.getElementById("from").value = document.getElementById("to").value;
  document.getElementById("to").value = tmp;
}

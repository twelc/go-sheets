

function applyFilters(range){
  q = $("#querry").val().toLowerCase()
  min = $("#min").val()
  max = $("#max").val()
  $.ajax({
    type: "post",
    url: "/get-table-filter",
    data: {"querry":q, "min":min, "max":max, "range":range},
    dataType: "html",
    success: function (response) {
      $("#main").html(response)
    }
  });
  post = 1
  filtered = true
}


$("#send").click(function(){applyFilters(0)})

function throttle(callee, timeout) {
  let timer = null

  return function perform(...args) {
    if (timer) return

    timer = setTimeout(() => {
      callee(...args)

      clearTimeout(timer)
      timer = null
    }, timeout)
  }
}

let post = 1
let isLoading = false
let filtered = false

function checkPosition() {
  // Нам потребуется знать высоту документа и высоту экрана:
  const height = document.body.offsetHeight
  const screenHeight = window.innerHeight

  // Они могут отличаться: если на странице много контента,
  // высота документа будет больше высоты экрана (отсюда и скролл).

  // Записываем, сколько пикселей пользователь уже проскроллил:
  const scrolled = window.scrollY

  // Обозначим порог, по приближении к которому
  // будем вызывать какое-то действие.
  // В нашем случае — четверть экрана до конца страницы:
  const threshold = height - screenHeight

  // Отслеживаем, где находится низ экрана относительно страницы:
  const position = scrolled + screenHeight

  if (position >= threshold) {
    if (post == -1){
      return
    }
    if (isLoading) {
      return
    }
    if (filtered){
      applyFilters(post)
    }
    isLoading = true
    $.ajax({
      type: "post",
      url: "/get-table-data",
      data: {"range":post},
      dataType: "html",
      success: function (response) {
        if (!response){
          post = -1
          isLoading = false
          return
        }

        $("#main").append(response)
        post++
        isLoading = false
      }
    });
  }
}

;(() => {
  window.addEventListener('scroll', throttle(checkPosition, 500))
  window.addEventListener('resize', throttle(checkPosition, 500))
})()

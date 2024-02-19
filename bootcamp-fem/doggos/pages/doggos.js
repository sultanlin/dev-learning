// const BREEDS_URL = 'https://dog.ceo/api/breeds/image/random'
//
// function addDoggo() {
//   /* let dogImg = await */
//   fetch(BREEDS_URL)
//     .then(function (response) {
//       return response.json()
//     })
//     .then(function (data) {
//       const img = document.querySelector('img')
//       img.src = data.message
//       img.alt = 'Cute doggo'
//       // console.log(data)
//       document.querySelector('.doggos').appendChild(img)
//     })
//
//   console.log('which will log first?')
// }
//
// document.querySelector('.add-doggo').addEventListener('click', addDoggo)

const BREEDS_URL = 'https://dog.ceo/api/breeds/list/all'
const img = document.querySelector('img')
const select = document.querySelector('.breeds')

function addDoggo() {
  // Get all dogs from api
  // Add each dog as an option
  fetch(BREEDS_URL)
    .then(function(response) {
      return response.json()
    })
    .then(function(data) {
      const dogJson = data.message
      const dogList = Object.keys(dogJson)
      console.log(dogList)
      for (let dog of dogList) {
        const option = document.createElement('option')
        option.value = dog
        option.innerText = dog
        // option.innerText(dog)
        select.appendChild(option)
      }
    })
}
function getDoggo(url) {
  fetch(url)
    .then((response) => response.json())
    .then((data) => {
      img.src = data.message
    })
}

function init() {
  addDoggo()

  select.addEventListener('change', (event) => {
    let url = `https://dog.ceo/api/breed/${event.target.value}/images/random`
    getDoggo(url)
  })
}

init()

import "./styles/scss/main.scss";
import { Selectable } from './jsm/Selectable';



//let attention = Prompt();
const dropdown = document.querySelector('.js-nav');

dropdown.addEventListener('click', function(e) {
    const dropDownTarget = e.target.closest('.nav-item');


    if (e.target.id === 'dropDownMenuButton') {
        // Hide the dropdown
        dropDownTarget.nextElementSibling.classList.toggle('hidden');

        if (!dropDownTarget.nextElementSibling.classList.contains('hidden')) {
            e.target.style.color = '#181614';
        }

        // Toggle aria expanded attributes
        ariaToggle(e.target)

        // Click outside
        hideOnClickOutside(e.target);

        e.preventDefault();
    }

})

const hideOnClickOutside = (element) => {
    const outsideClickListener = (event) => {
        if (!element.contains(event.target)) {
            if (!!element && !!(element.offsetWidth || element.offsetHeight || element.getClientRects().length)) {
                element.nextElementSibling.classList.toggle('hidden');
                removeClickListener();
                element.style.color = '';

            }
        }
    }
    const removeClickListener = () => {
        document.removeEventListener('click', outsideClickListener);
    }
    document.addEventListener('click', outsideClickListener);
}



const ariaToggle = (toggle) => {
    toggle.getAttribute('aria-expanded') === 'false' ?
        toggle.setAttribute('aria-expanded', 'true') :
        toggle.setAttribute('aria-expanded', 'false');
}


const toggleNav = () => {
    const toggleNav = document.querySelector('.nav-brand-toggle');

    toggleNav.addEventListener('click', function (e) {
        const navContentWrapper = document.querySelector('.navContentWrapper');
        navContentWrapper.classList.toggle('hidden')
    })
}

toggleNav();


// Validate form
class FormValidator {
    constructor(form, fields) {
        this.form = form;
        this.fields = fields;
    }

    validateOnSubmit() {
        let self = this;
        this.form.addEventListener('submit', (e) => {
            e.preventDefault();
            Array.from(self.fields, field =>{
                self.validateFields(field);
            })
        })

    }

    validateOnEntry() {
        let self = this;
        Array.from(self.fields, field =>{
            const input = field.getAttribute('id')
            const inputId = document.querySelector(`#${input}`)
            inputId.addEventListener('input', e => {
                self.validateFields(field);

            })


        })
    }

    validateFields(field) {
        if (field.value.trim() === '') {
            this.setStatus(field, `${field.previousElementSibling.innerHTML} cannot be blank`, "error")
        } else {
            this.setStatus(field, null, "success")
        }
        if (field.type === 'email') {
            const re = /^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
            if (re.test(field.value)) {
                this.setStatus(field, null, "success")
            } else {
                this.setStatus(field, `Please enter a valid email address`, "error")
            }
        }
        if (field.id === 'passwordConfirmation') {
            const passwordField = this.form.querySelector("#password");

            if (field.value.trim() === "") {
                this.setStatus(field, `Password confirmation required`, "error")
            } else if (field.value !== passwordField.value) {
                this.setStatus(field, `Password does not match`, "error")
            } else {
                this.setStatus(field, null, "success")
            }
        }

        if (field.id === 'endDate') {
            const startDate = this.form.querySelector('#startDate');

            const start = new Date(startDate.value)
            const end = new Date(field.value)

            if (start.getTime() >= end.getTime()) {
                this.setStatus(field, `Select date after your arrival`, "error")
            } else {
                this.setStatus(field, null, "success")
            }

        }
    }

    setStatus(field, message, status) {
        const errorMessage = field.parentElement.querySelector('.error-message');
        const formAlert = field.parentElement.querySelector('.form-alert');

        if (status === 'success') {
            formAlert.classList.remove('is-danger')
            formAlert.classList.add('is-success')
            field.classList.remove('is-danger')
            field.classList.add('is-success')
            if (errorMessage) { errorMessage.innerHTML = '' }

        }

        if (status === 'error') {
            formAlert.classList.remove('is-success')
            errorMessage.innerHTML = message
            formAlert.classList.add('is-danger')
            field.classList.add('is-danger')
        }
    }


    init() {
        this.validateOnSubmit()
        this.validateOnEntry()
    }
}

const form = document.querySelector('.js-validateForm');
const fields = document.querySelectorAll('.js-validateForm input')
const validator = new FormValidator(form, fields)

if (form != null) {
    validator.init();
}


/* Custom select */
const labels = document.querySelector('.custom-labels');
const genres = document.querySelector('.custom-genres');
const artists = document.querySelector('.custom-artists');
const trackArtists = document.querySelector('.custom-track-artists');


if (labels) {
    new Selectable(labels).init();
}

if (genres) {
    new Selectable(genres).init();
}

if (artists) {
    new Selectable(artists).init();
}

if (trackArtists) {
    new Selectable(trackArtists).init();
}



// var regx = /[0-9 ()-]/;
// var inp = document.getElementById('field1');

// inp.addEventListener('keypress', function(e) {
//     console.log('key: ' + e.key);
//     if (! e.key.match(regx)) {
//         // if the key doesn't match the allowed characters,
//         // prevent it from being entered.
//         e.preventDefault();
//     }
// });




// const customSelects = document.querySelectorAll('.custom-select');


// customSelects.forEach(customSelect => {
//     customSelect.addEventListener('click', (e) => {
//         e.currentTarget.parentElement.classList.toggle('show');

//         const option = e.currentTarget.nextElementSibling;
//         const options = option.querySelectorAll('.custom-option li');
//         const customOption = option.querySelector('.custom-option');

//         const value = customSelect.querySelector('.value')
//         const inputCategory = customSelect.querySelector('#inputCategory');




//         options.forEach(item=> {
//             item.addEventListener('click', function (e) {
//                 value.innerHTML = e.target.innerHTML;
//                 inputCategory.options[0].value = e.target.dataset.id;
//                 e.target.parentElement.classList.remove('show');
//             })


//         })

//         // const unique = [...new Set(selectedOptions)]
//         //
//         //
//         // for (var i = 0; i <= unique.length; i++) {
//         //
//         //     var input = document.createElement('input');
//         //     input.type = 'hidden';
//         //     input.name = 'inputCategories';
//         //     input.value = unique[i];
//         //
//         //
//         //     // var inputConatiner = document.createElement('div');
//         //     //
//         //     // inputConatiner.insertAdjacentHTML("afterbegin",input)
//         //     //
//         //     if (input.value != 'undefined') {
//         //         const selectForm = document.querySelector('.select');
//         //         selectForm.appendChild(input)
//         //     }
//         //
//         //     //
//         //     // selectForm.insertAdjacentHTML("afterbegin",  inputConatiner)
//         //
//         // }






//         window.addEventListener('click', function (e) {
//             if(e.target !== customSelect && e.target !== customOption) {
//                 customSelect.parentElement.classList.remove('show');
//             }
//         })

//     })
// })
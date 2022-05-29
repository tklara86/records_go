import "./styles/scss/main.scss";



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
class Selectable {

    #options;
    #input;
    #customOptionsSelectedDiv;

    constructor (element) {
        this.element = element;
        this.#options = this.element.nextElementSibling;
        this.#customOptionsSelectedDiv = this.#options.nextElementSibling;
        this.#input = this.#options.querySelector('input');
    }

    toggleSelect = () => {
        this.element.parentElement.classList.toggle('show');  

    }

    closeSelect = (e) => {
        var option = e.target.closest('.option') 
        if (e.target != this.element && option === null) {
            this.element.parentElement.classList.remove('show');
        }  
    }

    searchSelect = () => {
    
        this.#input.addEventListener('keyup', (e) => {
            let filteredName = e.target.value;

            const jsonResults = this.#input.nextElementSibling;
            const labels = jsonResults.querySelectorAll('label');
            const regex = new RegExp(filteredName, "gi");
            const noFound = jsonResults.querySelector('.no-results-create');
            
            labels.forEach(label => {
                let labelInput = label.querySelector('input');
             
                
                labelInput.dataset.input.match(regex) 
                    ? label.style.display = 'flex' 
                    : label.style.display = 'none'; 


                                
                    if (!labelInput.dataset.input.match(regex)) {
                        if (noFound) {
                            noFound.remove();
                        
                        } 
                    }

            });

            let m = `
                <div class="no-results-create">
                    <div class="no-results-found">
                        <span class="icon exclamation-icon"></span>
                        <p>No results found!</p>
                    </div>
                    <div class="add-to-db">
                        <span id="js-inputVal">Would you like to add <span class="input-value"></span></span>
                        <div id="js-buttonAdd">
                        <svg class="add-icon js-add-icon">
                            <use xlink:href="/dist/icons/sprite.svg#file"></use>
                        </svg>
                        <p>Add</p>
                        </div>
                    </div>
                </div>
            `;

           

            if (jsonResults.innerText === '') {
                jsonResults.insertAdjacentHTML('beforeend', m);
                const inputValue = document.querySelector('.input-value');

                let attachedEvent = false;
                const buttonAdd = jsonResults.querySelector('#js-buttonAdd');

                if (!attachedEvent) {
                    buttonAdd.addEventListener('click', () => {
                        switch (jsonResults.id) {
                            case 'labels':
                                
                                try {
                                    const data = this.postData('http://localhost:8000/admin/record/postLabelsJSON', { 
                                       name: filteredName
                                    })
                                
                                } catch (error) {
                                    console.log(error) 
                                }
                           
                                break;
                            case 'genres':
                                try {
                                    const data = this.postData('http://localhost:8000/admin/record/postGenresJSON', { 
                                       name: filteredName
                                    })
                                    
                        
                                } catch (error) {
                                    console.log(error) 
                                }
                
                                break;
                            case 'artists':
                                try {
                                    const data = this.postData('http://localhost:8000/admin/record/postArtistsJSON', { 
                                       name: filteredName
                                    })
                                    
                        
                                } catch (error) {
                                    console.log(error) 
                                }
                                break;
                        }
            
                        attachedEvent = true;
                    })
                }
      
                inputValue.innerHTML = filteredName;  
            }  

        });
    }
    postData = (url, data) => {
       const res =  fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        }).then(res => res.json())

        return res;
    
    }

    showSelectedResults = () => {
        const jsonResults = this.#input.nextElementSibling;
    
        jsonResults.addEventListener('change', (e) => {
            const input = e.target.closest('input');

            if (input.checked) {
                let markup = `
                <span class="custom-options-selected--item" data-index="${input.value}">${input.dataset.input}
                    <span class="icon-close"></span>
                </span>
                `;  
                
                this.#customOptionsSelectedDiv.insertAdjacentHTML('afterbegin', markup); 

                let customOptions = this.#customOptionsSelectedDiv.querySelectorAll('.custom-options-selected--item');

                customOptions.forEach(co => {
                    let icon = co.querySelector('.icon-close');
                    icon.addEventListener('click', (e) => {
                        let index2 = e.target.parentElement.dataset.index;
                        let parentEl = e.target.parentElement
                        parentEl.remove();

                        if (input.value == index2) {
                            input.checked = false;
                        }
                    })
                })
            } else {
                let customOptionsAll = this.#customOptionsSelectedDiv.querySelectorAll('.custom-options-selected--item');
                customOptionsAll.forEach(customOption => {
                    if (input.value == customOption.dataset.index) {
                        customOption.remove();
                    }  
                })
            }

        })
    }

    init = () =>  {
        // toggle 'select'
        this.element.addEventListener('click', this.toggleSelect);
        // close 'select' when click outside container
        window.addEventListener('click', this.closeSelect);
        // search inputs
        this.searchSelect();
        this.showSelectedResults();
    }
}




const labels = document.querySelector('.custom-labels');
const genres = document.querySelector('.custom-genres');
const artists = document.querySelector('.custom-artists');


if (labels) {
    new Selectable(labels).init();
}

if (genres) {
    new Selectable(genres).init();
}

if (artists) {
    new Selectable(artists).init();
}

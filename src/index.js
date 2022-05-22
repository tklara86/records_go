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
const customSelects = document.querySelectorAll('.custom-select');
let outputs = [];
let outputs2 = [];
let ajaxOutputs = [];
let inputsAll;



window.addEventListener('load', () => {

    let jsonResults = document.querySelectorAll('.json-results');
    let markup;

    fetch('http://localhost:8000/admin/record/labelsJSON')
    .then(res => res.json())
    .then(data => {
        for (var j = 0; j < data.labels.length; j++) {
            ajaxOutputs.push(data.labels[j])
            const { id, name, input_name } = data.labels[j];
            markup = `
                <label>
                    <input type="checkbox" name="${input_name}" value="${id}" data-input="${name}">${name}
                </label>
            `;

            for (var i = 0; i < jsonResults.length; i++) {
                if (jsonResults[i].id == 'labels') {
                    jsonResults[i].insertAdjacentHTML('afterbegin', markup);
                }
                
            }
        } 
    });   
});

customSelects.forEach(customSelect => {
    
    // Toggle selects
    customSelect.addEventListener('click', (e) => {
        e.currentTarget.parentElement.classList.toggle('show');

            // Check input
        const options = customSelect.nextElementSibling;
        const customOptionsSelectedDiv = options.nextElementSibling;
        const inputs = options.querySelectorAll('input');
        let markup;
    

        if (e.currentTarget.parentElement.classList.contains('show')) {
            inputs.forEach((input,index) => {



                const allInputs = [];
                var customOption = options.querySelector('.custom-option');
                input.addEventListener('keyup', (e) => {
                    let filteredName = e.target.value;
                    
                    const jsonResults = e.target.nextElementSibling;
                    const labels = jsonResults.querySelectorAll('label');
    
            
                    ajaxOutputs.filter(object => {
                        const regex = new RegExp(filteredName, "gi")
                        if (object.name.match(regex)) {
                            labels.forEach(label => label.remove());
    
                            let markup = `
                                <label>
                                 <input type="checkbox" name="${object.input_name}" value="${object.id}" data-input="${object.name}">${object.name}
                                </label> 
                            `
                            jsonResults.insertAdjacentHTML('afterbegin', markup)
    
                            let inputsAttached = jsonResults.querySelectorAll('input');
    
                            inputsAttached.forEach(inputAt => {
                                inputAt.addEventListener('change', (e) => {
                                    if (e.target.checked) {
                                        outputs2.push(inputAt)
    
                                 
                                        outputs2.forEach((output) => {
                                            markup = `
                                                <span class="custom-options-selected--item" data-index="${index}">${output}
                                                    <span class="icon-close"></span>
                                                </span>
                                            `;
                                        });
                        
                                        customOptionsSelectedDiv.insertAdjacentHTML('afterbegin', markup); 
                                     
                                        
                                    }
                                    
                                })
                            })
                            
                        }
                      
                    })
                  
    
                })
    
    
               
          
            input.addEventListener('change', (e) => {
                if (e.target.checked) {
                    outputs.push(input.dataset.input);
                    
    
                    outputs.forEach((output) => {
                        markup = `
                            <span class="custom-options-selected--item" data-index="${index}">${output}
                                <span class="icon-close"></span>
                            </span>
                        `;
                    });
    
                 
                 
                    customOptionsSelectedDiv.insertAdjacentHTML('afterbegin', markup);
                    let customOptions = customOptionsSelectedDiv.querySelectorAll('.custom-options-selected--item');
                    customOptions.forEach(co => {
                        let icon = co.querySelector('.icon-close');
                        icon.addEventListener('click', (e) => {
                            let index2 = e.target.parentElement.dataset.index;
                            let parentEl = e.target.parentElement
                            parentEl.remove();
    
                            if (index2 == index) {
                                input.checked = false;
                            }
                        })
                    })
               
                    
                } else {
                   // outputs.splice(outputs.indexOf(input.dataset.input), 1);
                   let customOptionsAll = customOptionsSelectedDiv.querySelectorAll('.custom-options-selected--item');
                    customOptionsAll.forEach(customOption => {
                        if (index == customOption.dataset.index) {
                            customOption.remove();
                        }
                        
                    })
                  
                  
                }
            });
    
    
    
        }); 
        }
        
    });

    // Click outside select
    window.addEventListener('click', function (e) {
        var option = e.target.closest('.option') 
        if (e.target != customSelect && option === null) {
            customSelect.parentElement.classList.remove('show');
        }
    });


  

    

})


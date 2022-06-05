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
const addTracks = document.querySelector('.custom-tracks');


if (labels) {
    new Selectable(labels).initMultiple();
}

if (genres) {
    new Selectable(genres).initMultiple();
}

if (artists) {
    new Selectable(artists).initMultiple();
}



if (addTracks) {
    new Selectable(addTracks).initSingle();  
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



const addTrackButton = document.getElementById('addTrackButton');



function addMultipleTracks() {
    const countTracks = document.getElementById('countTracks');
    const numberOfTracks = parseInt(countTracks.innerHTML);

    const tracklistTableHead = document.querySelector('.tracklist-table-head');

    for (var i = 0; i < numberOfTracks; i++) {

        
        let markup = `
            <tr>
            <td class="td-width-sm">  
            <div class="form-control">
                <span class="form-alert"></span>
                <input autocomplete="off" class="input-control input-control--small" type="text" name="record-track" placeholder="Track">
            </div>
            </td>
            <td class="td-flex">
            <div class="add-tracklist-artist-container">
                <div class="add-track-artist">
                <svg class="add-icon small">
                    <use xlink:href="/dist/icons/sprite.svg#plus-circled"></use>
                </svg>
                Add Artist
                </div>
            </div>
            
        
            </td>
            <td>
            <div class="form-control">
                <span class="form-alert"></span>
                <input autocomplete="off" class="input-control input-control--small" type="text" name="record-track-title" placeholder="Track Title">
            </div>
            </td>
            <td class="td-width-md">
            <div class="form-control">
                <span class="form-alert"></span>
                <input autocomplete="off" class="input-control input-control--small" type="text" name="record-track-duration" placeholder="0:00" maxlength="3">
            </div>
            </td>
            <td class="td-flex td-width-sm">
            <div class="remove icon-flex removeTrack">
                <svg class="action-icon sm">
                <use xlink:href="/dist/icons/sprite.svg#trash"></use>
                </svg>
                Remove
            </div>
            </td>
            </tr>
            `;

            tracklistTableHead.insertAdjacentHTML('beforeend', markup);
    }
  

    const removeTracks = document.querySelectorAll('.removeTrack');

    removeTracks.forEach(removeTrack => {
        removeTrack.addEventListener('click', (e) => {
            e.target.parentElement.parentElement.remove();
        }) 
    })



    const addTrackArtists = document.querySelectorAll('.add-track-artist');

    addTrackArtists.forEach(addTrackArtist => {
        let isAdded = false;
        addTrackArtist.addEventListener('click', (e) => {
            const parentEl = e.target.parentElement;
            const tr = parentEl.parentElement;
            
            
            let m = `
                    <div class="form-control sub-select">
                    <div id="artists" class="select">
                    <div class="custom-select custom-track-artists">
                        <span class="value">Select artist</span>
                        <span class="icon dropdown-icon"></span>
                    </div>
                    <div class="option">
                        <div class="custom-option">
                        <input autocomplete="off" class="input-control input-control--small input-control--nested filter" type="text" name="search" placeholder="Search">
                        <div id="track-artists" class="json-results">
                        </div>
                        </div>
                    </div>
                    <div class="custom-options-selected"></div>
                    </div> 
                    <div class="remove icon-flex removeTrackArtist">
                        <svg class="action-icon sm">
                        <use xlink:href="/dist/icons/sprite.svg#trash"></use>
                        </svg>
                        Remove
                    </div>
                </div>
            `;


            if (!isAdded) {
                parentEl.insertAdjacentHTML('afterend', m);
                isAdded = true;
            }

            const trackArtist = tr.querySelector('.custom-track-artists');
      
            new Selectable(trackArtist).initMultiple();
            
            const options = parentEl.nextElementSibling;
            const jsResults = options.querySelector('#track-artists');

          
          
            fetch('http://localhost:8000/admin/record/artistsJSON')
                .then(res => res.json())
                .then(data => {  
     
                for (var j = 0; j < data.artists.length; j++) {    
                    const { id, name, input_name } = data.artists[j];
                    let markup = `
                        <label>
                            <input type="checkbox" name="${input_name}-track" value="${id}" data-input="${name}">${name}
                        </label>
                    `;

                    jsResults.insertAdjacentHTML('afterbegin', markup);
                    
                } 
            });


            const removeTrackArtist = tr.querySelector('.removeTrackArtist');

            if (removeTrackArtist) {
                removeTrackArtist.addEventListener('click', (e) => {
                   e.target.parentElement.remove();
                   isAdded = false
                  
                })
            }

            
        
        })



 
    })
}


addTrackButton.addEventListener('click', addMultipleTracks);







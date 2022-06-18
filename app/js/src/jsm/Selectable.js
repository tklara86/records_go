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
        
            
            if (labelInput.dataset.input.match(regex)) {
                label.style.display = 'flex';
                if (noFound) {
                    noFound.remove();
                } 
            }  else {
                label.style.display = 'none';  
            }

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
                                }, 'http://localhost:8000/admin/record/labelsJSON', jsonResults)

                                if (data) {
                                    this.#input.value = '';
                                    const noFound2 = jsonResults.querySelector('.no-results-create');
                                    if (noFound2) {
                                        noFound2.remove()
                                    }

                                    labels.forEach(label => label.remove());

                                
                                }

                            
                            
                            } catch (error) {
                                console.log(error) 
                            }
                            break;
                        case 'genres':
                            try {
                                const data = this.postData('http://localhost:8000/admin/record/postGenresJSON', { 
                                    name: filteredName
                                }, 'http://localhost:8000/admin/record/genresJSON', jsonResults)
                                if (data) {
                                    this.#input.value = '';
                                    const noFound2 = jsonResults.querySelector('.no-results-create');
                                    if (noFound2) {
                                        noFound2.remove()
                                    }
                                    labels.forEach(label => label.remove());
                                }
                                
                    
                            } catch (error) {
                                console.log(error) 
                            }
            
                            break;
                        case 'artists':
                            try {
                                const data = this.postData('http://localhost:8000/admin/record/postArtistsJSON', { 
                                    name: filteredName
                                }, 'http://localhost:8000/admin/record/artistsJSON', jsonResults)
                                if (data) {
                                    this.#input.value = '';
                                    const noFound2 = jsonResults.querySelector('.no-results-create');
                                    if (noFound2) {
                                        noFound2.remove()
                                    }
                                    labels.forEach(label => label.remove());
                                }
                                
                    
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
  postData = (url, data, getUrl, results) => {
     const res =  fetch(url, {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json'
          },
          body: JSON.stringify(data)
      }).then(res => res.json()).then(() => this.getData(getUrl, results))

      return res;
  
  }


  getData = (url, results) => {
      fetch(url)
          .then(res => res.json())
          .then(data => {  
              let dataObj;

              if (data.labels) {
                  dataObj = data.labels;
              } else if (data.artists) {
                  dataObj = data.artists
              } else if (data.genres) {
                  dataObj = data.genres
              }

  
              for (var j = 0; j < dataObj.length; j++) {    
                  const { id, name, input_name } = dataObj[j];
                  let markup = `
                      <label>
                          <input type="checkbox" name="${input_name}" value="${id}" data-input="${name}">${name}
                      </label>
                  `;

                  results.insertAdjacentHTML('afterbegin', markup);
                  
              } 
          });
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


  selectValue = () => {

        const options = document.querySelectorAll('.custom-option li');
        const value = this.element.querySelector('.value');
        const parentEl = this.element;

        options.forEach(item=> {
            item.addEventListener('click', function (e) {
                value.innerHTML = e.target.innerHTML;
                parentEl.parentElement.classList.remove('show');

            })
        })
  }

  initMultiple = () =>  {
      // toggle 'select'
      this.element.addEventListener('click', this.toggleSelect);
      // close 'select' when click outside container
      window.addEventListener('click', this.closeSelect);
      // search inputs
      this.searchSelect();
      this.showSelectedResults();
  }

  initSingle = () => {
     // toggle 'select'
     this.element.addEventListener('click', this.toggleSelect);
     // close 'select' when click outside container
     window.addEventListener('click', this.closeSelect);  
     this.selectValue();
  }
}



export { Selectable }
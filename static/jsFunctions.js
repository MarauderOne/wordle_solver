let activeBox = null;

document.addEventListener("DOMContentLoaded", function() {
    // Focus the first box by default
    const firstBox = document.querySelector('.box');
    if (firstBox) {
        firstBox.focus();
        activeBox = firstBox;
    }

    // Track the last focused input box
    document.querySelectorAll('.box').forEach(box => {
        box.addEventListener('focus', (event) => {
            activeBox = event.target;
        });
    });

    // Function to handle key press
    function handleKeyPress(event) {
        event.preventDefault();  // Prevent the button from stealing focus
        const character = event.target.getAttribute('character').toUpperCase();
        if (activeBox) {
            activeBox.value = character;
            // Optionally, move focus to the next box
            const nextBox = activeBox.nextElementSibling;
            if (nextBox && nextBox.classList.contains('box')) {
                nextBox.disabled = false;
                nextBox.focus();
            }
        }
    }

    // Function to handle delete key press
    function handleDeleteKeyPress(event) {
        event.preventDefault();  // Prevent default behavior
        if (activeBox) {
            activeBox.value = '';  // Clear the contents of the active box
            // Move focus to the previous box if it exists
            const previousBox = activeBox.previousElementSibling;
            if (previousBox && previousBox.classList.contains('box')) {
                previousBox.focus();
            }
        }
    }


    // Add event listeners to each keyboard button
    const keyboardKeys = document.querySelectorAll('.keyboardKey:not(.keyboardKey[character="←"])');
    keyboardKeys.forEach(key => {
        key.addEventListener('click', handleKeyPress);
    });

    // Add event listener to the delete key button
    const deleteKey = document.querySelector('.keyboardKey[character="←"]');
    deleteKey.addEventListener('click', handleDeleteKeyPress);

    function limitInputToSingleChar(event) {
        const input = event.target;
        if (input.value.length > 1) {
            input.value = input.value.charAt(0);
        }
        // Enable the next box if this one is filled
        if (input.value.length === 1) {
            enableNextBox(input);
        }
        sendGridState();
    }

    function enableNextBox(currentBox) {
        const boxes = Array.from(document.querySelectorAll('.box'));
        const currentIndex = boxes.indexOf(currentBox);
        if (currentIndex < boxes.length - 1) {
            boxes[currentIndex + 1].disabled = false;
            boxes[currentIndex + 1].focus();
        }
    }

    function changeBoxColor(box) {
        if (box.value.trim() === "") {
            return; // Do nothing if the box is empty
        }
    
        const currentClass = box.className.split(' ').find(cls => ['grey', 'yellow', 'green'].includes(cls));
        let newClass = 'grey';
        if (currentClass === 'grey') newClass = 'yellow';
        if (currentClass === 'yellow') newClass = 'green';
        if (currentClass === 'green') newClass = 'grey';
        box.classList.remove('grey', 'yellow', 'green');
        if (newClass) box.classList.add(newClass);
        sendGridState();
    }

    async function sendGridState() {
        const boxes = document.querySelectorAll('.box');
        const gridData = Array.from(boxes).map(box => ({
            character: box.value.toUpperCase(),
            color: box.className.split(' ').find(cls => ['grey', 'yellow', 'green'].includes(cls)) || ''
        }));
    
        try {
            const response = await fetch('/guesses', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(gridData)
            });
            const data = await response.json();
            document.getElementById('resultsCount').innerText = data.resultSummary;
            document.getElementById('results').innerText = data.result;
            if (data.error) {
                throw document.getElementById('resultsCount').innerText = 0,
                document.getElementById('results').innerText = data.error;
            }
        } catch (error) {
            console.error('error:', error);   
        }
    }

    const boxes = document.querySelectorAll('.box');
    boxes.forEach(box => {
        box.addEventListener('input', limitInputToSingleChar);
        box.addEventListener('click', () => changeBoxColor(box));
    });
});
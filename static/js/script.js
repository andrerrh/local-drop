const matchAndSwapFiles = (files, target, regexp) => {
	const nameMatch = new RegExp(regexp, "i");
	const filesToDisplay = files.filter((file) => nameMatch.test(file.children[0].outerText));
	target.replaceChildren(...filesToDisplay);
}

document.addEventListener("DOMContentLoaded", () => {
	const dropzone = document;
	const fileInputField = document.querySelector("#up-files");

	["dragenter", "dragover"].forEach(event => {
		dropzone.addEventListener(event, e => {
			e.preventDefault();
			dropzone.body.classList.add("bg-gray-300");
		});
	});

	["dragleave", "drop"].forEach(event => {
		dropzone.addEventListener(event, e => {
			e.preventDefault();
			dropzone.body.classList.remove("bg-gray-300");
		});
	});

	dropzone.addEventListener("drop", e => {
		fileInputField.files = e.dataTransfer.files;
		fileInputField.dispatchEvent(new Event("change", { bubbles: true }));
	})

	//Handles part of the request response upload of files that already exist
	document.body.addEventListener("htmx:beforeSwap", (event) => {
		if (event.detail.xhr.status === 409) {
			event.detail.shouldSwap = true;
			event.detail.isError = false;
		}
	})
})


document.addEventListener("htmx:afterSwap", () => {
	const FilesContainer = document.querySelector("#files-container");
	const filesChildren = Array.from(FilesContainer.children);
	const fileSearchInput = document.querySelector("#search-input");

	fileSearchInput.addEventListener("keyup", () => {
		matchAndSwapFiles(filesChildren, FilesContainer, fileSearchInput.value);
	})
	document.querySelector("#up-btn").addEventListener("click", () => {
		document.querySelector("#up-files").click()
	})
})

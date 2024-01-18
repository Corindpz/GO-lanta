function showImageGallery() {
  var gallery = document.getElementById("image-gallery");
  var fileInput = document.getElementById("file-upload");

  // Afficher la galerie et masquer le bouton et le champ de fichier
  gallery.style.display = "flex";
  fileInput.style.display = "none";
}

function selectImage(selectedImage) {
  var preview = document.getElementById("selected-image-preview");
  var fileInput = document.getElementById("file-upload");
    console.log("a");
  // Afficher l'image sélectionnée
  while (preview.firstChild) {
    preview.removeChild(preview.firstChild);
  }
  console.log("b");
  var img = document.createElement("img");
  img.src = selectedImage.src;
  img.style.width = "100px";
  img.style.height = "100px";
  preview.appendChild(img);
  console.log("c");
  // Mettre à jour la valeur du champ de fichier avec le chemin de l'image sélectionnée
  fileInput.value = selectedImage.src;

  // Masquer la galerie après la sélection
  document.getElementById("image-gallery").style.display = "none";
}

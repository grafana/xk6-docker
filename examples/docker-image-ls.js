import images from 'k6/x/docker/images';

function logImage(image) {
  console.log(JSON.stringify(image, null, 2))
}

export default function () {
  const imagesList = images.list()
  console.log(`${imagesList.length} images found:`)
  imagesList.forEach(logImage)
}

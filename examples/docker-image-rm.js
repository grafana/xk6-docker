import images from 'k6/x/docker/images';

export default function () {
  const imageId = 'YOUR_IMAGE_ID'
  const result = images.remove(imageId, {force: true})
  console.log(`Image ${imageId}removal result: ${JSON.stringify(result)}`)
}

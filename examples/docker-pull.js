import images from 'k6/x/docker/images';

export default function () {
  const imageName = 'YOUR_IMAGE_NAME'
  const result = images.pull(imageName)
  console.log(`Image pulling result: ${result}`)
}

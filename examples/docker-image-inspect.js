import images from 'k6/x/docker/images';

export default function () {
  const imageId = 'YOUR_IMAGE_ID';
  const data = images.inspect(imageId);
  console.log(JSON.stringify(data, null, 2));
}

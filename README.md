# Photomarker

Photomarker is a simple CLI tool to add watermarks to photos. It allows to mark multiple photos by given size, location, and opacity.

Example usage to add the desired watermark to the photos in the current directory is as follows.
```shell
photomarker -dst=. -src=watermark.png -width=100 -height=100 -x=-50 -y=-50 
-opacity=0.3 -out=marked_photos/
```

## Installation
You can use the executable files under the builds directory by downloading the appropriate one for your system. If you want to access the tool globally you can add the executable to your working path.

If there is no suitable executable file for your system, you can obtain the executable file with the compiler of the Go language.

```shell
go get github.com/sschrs/photomarker
```
```shell
go build -o photomarker main.go # In directory of package 
```

## Usage and Arguments
To see usage of args you can use -help flag.
```shell
photomarker -help
```

<table>
  <thead>
    <tr>
      <th>Argument</th>
      <th>Usage</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>width</td>
      <td>Specifies the width, in pixels, of the photo to be placed. Default: 60px</td>
    </tr>
    <tr>
        <td>height</td>
        <td>Specifies the height, in pixels, of the photo to be placed. Default: 60px</td>
    </tr>
    <tr>
        <td>opacity</td>
        <td>Opacity value of the photo to be placed. Min: 0, Max:1, Default: 0.5</td>
    </tr>
    <tr>
        <td>out</td>
        <td>If the replace argument is false, it specifies the destination path of the copied photos. Default: marked/</td>
    </tr>
    <tr>
        <td>x</td>
        <td>Specifies the x coordinate of the photo to be placed. If a negative value is entered, pixels from the lower right corner are counted. Default: -20</td>
    </tr>
    <tr>
        <td>y</td>
        <td>Specifies the y coordinate of the photo to be placed. If a negative value is entered, pixels from the lower right corner are counted. Default: -20</td>
    </tr>
    <tr>
        <td>dst</td>
        <td>Destination file or directory</td>
    </tr>
    <tr>
        <td>src</td>
        <td>Source file to be used for markup</td>
    </tr>
  </tbody>
</table>

## TODO
* Specify the output format
* Replace feature

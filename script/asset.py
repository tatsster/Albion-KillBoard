# Quick script to download all assets
import requests
import os
import sys
import concurrent.futures

class BaseURL:
    def __init__(self, url: str):
        self.url = url

    def BuildURL(self, path: str, path_params: dict, query: dict) -> str:
        # Replace path params in the path string
        for param_name, param_value in path_params.items():
            path = path.replace(f":{param_name}", param_value)

        # Build the URL with path and query parameters
        url = path
        if query:
            query_params = "&".join(f"{key}={value}" for key, value in query.items())
            url += "?" + query_params

        return "{}/{}".format(self.url, url)

def main():
    # https://render.albiononline.com/v1/item/T8_2H_DUALSCIMITAR_UNDEAD@4.png?count=1&quality=4&size=128
    items_path = "assets/items.txt"
    asset_url = BaseURL("https://render.albiononline.com/v1/item")
    os.makedirs("assets/items", exist_ok=True)

    startID = 1
    if len(sys.argv) > 1:
        startID = int(sys.argv[1])

    result_url = {}
    try:
        with open(items_path, "r") as file:
            for line in file:
                array = line.strip().split(": ")
                stt = int(array[0])
                if stt < startID:
                    continue

                array[1] = array[1].strip()
                if stt < 1642 or stt > 8481:
                    url = asset_url.BuildURL(":item_id.png", 
                                {
                                    "item_id": array[1]
                                },
                                {
                                    "quality": 0,
                                    "size": 128,
                                })
                    result_url["{}_{}".format(array[1], 0)] = url
                else:
                    for i in range(1, 6):
                        url = asset_url.BuildURL(":item_id.png", 
                                {
                                    "item_id": array[1]
                                },
                                {
                                    "quality": i,
                                    "size": 128,
                                })
                    result_url["{}_{}".format(array[1], i)] = url
    except FileNotFoundError:
        print(f"File '{items_path}' not found.")
    
    SaveImagesConcurrent(result_url)

def SaveImagesConcurrent(image_urls: dict):
    with concurrent.futures.ThreadPoolExecutor(max_workers=5) as executor:
        # Create a dictionary to map each download future to its corresponding image URL
        future_to_url = {executor.submit(SaveImage, url, name): url for name, url in image_urls.items()}

        # Iterate over completed futures as they become available
        for future in concurrent.futures.as_completed(future_to_url):
            url = future_to_url[future]
            try:
                result = future.result()
                print(result)
            except Exception as e:
                print(f"Error downloading image from {url}: {e}")

def SaveImage(url: str, fileName: str):
    response = requests.get(url)
    response.raise_for_status()  # Raise an exception if the request was not successful

    # Extract the file name from the URL
    file_name = "assets/items/{}.png".format(fileName)

    # Save the image to the specified path
    with open(file_name, 'wb') as file:
        file.write(response.content)
    print(f"Image downloaded successfully: {file_name}")

if __name__ == "__main__":
    main()
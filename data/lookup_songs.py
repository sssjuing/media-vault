
import csv
import requests
import urllib.parse

def search_song(title, artist=None):
    """Search song using iTunes Search API (free, no API key required)"""
    base_url = "https://itunes.apple.com/search?"
    
    terms = []
    if title:
        terms.append(title)
    if artist:
        terms.append(artist)
    
    if not terms:
        return None, None, None
    
    query = ' '.join(terms)
    params = {
        'term': query,
        'media': 'music',
        'entity': 'song',
        'limit': 1
    }
    
    url = base_url + urllib.parse.urlencode(params)
    
    try:
        response = requests.get(url, timeout=5)
        if response.status_code == 200:
            data = response.json()
            if data['resultCount'] > 0:
                result = data['results'][0]
                found_artist = result.get('artistName', '')
                found_album = result.get('collectionName', '')
                found_title = result.get('trackName', '')
                found_year = result.get('releaseDate', '').split('-')[0] if result.get('releaseDate') else ''
                return found_artist, found_album, found_title, found_year
        return None
    except Exception as e:
        print(f"Error searching for {query}: {e}")
        return None

def main():
    input_file = 'output_completed.csv'
    output_file = 'output_lookedup.csv'
    updated_count = 0
    
    with open(input_file, 'r', encoding='utf-8', errors='ignore') as infile:
        reader = csv.DictReader(infile)
        fieldnames = reader.fieldnames
        if not fieldnames:
            print("No fieldnames found")
            return
        
        rows = []
        for row in reader:
            need_update = False
            
            # Check if we have empty artist and title is filled
            title = row['title'].strip()
            current_artist = row['artist'].strip()
            current_album = row['album'].strip()
            current_year = row['year'].strip()
            
            if (not current_artist or not current_album) and title:
                print(f"Searching: {title} {current_artist}".strip())
                result = search_song(title, current_artist if current_artist else None)
                if result:
                    found_artist, found_album, found_title, found_year = result
                    if not current_artist and found_artist:
                        row['artist'] = found_artist
                        need_update = True
                    if not current_album and found_album:
                        row['album'] = found_album
                        need_update = True
                    if not current_year and found_year:
                        row['year'] = found_year
                    if need_update:
                        updated_count += 1
                        print(f"  Updated: Artist='{row['artist']}', Album='{row['album']}'")
            
            rows.append(row)
    
    with open(output_file, 'w', encoding='utf-8', newline='') as outfile:
        writer = csv.DictWriter(outfile, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(rows)
    
    print(f"\nLookup completed. Updated {updated_count} entries. Saved to {output_file}")

if __name__ == '__main__':
    main()

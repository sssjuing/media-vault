
import csv
import re
from urllib.parse import unquote

def parse_filename(filename):
    # Remove file extension
    name_without_ext = '.'.join(filename.split('.')[:-1])
    extension = filename.split('.')[-1].lower() if '.' in filename else ''
    
    album = ''
    artist = ''
    title = ''
    
    # Pattern 1: title_artist_album.mp3
    if '_' in name_without_ext:
        parts = name_without_ext.split('_')
        if len(parts) >= 3:
            title = parts[0].strip()
            artist = parts[1].strip()
            album = parts[2].strip()
        elif len(parts) == 2:
            # Could be title_artist or artist_title
            # Try to guess based on common patterns
            possible_title = parts[0].strip()
            possible_artist = parts[1].strip()
            # If first part looks like a title, second is artist
            title = possible_title
            artist = possible_artist
        elif len(parts) == 1:
            title = parts[0].strip()
    # Pattern 2: artist - title.mp3 or artist-title.mp3
    elif ' - ' in name_without_ext:
        parts = name_without_ext.split(' - ', 1)
        artist = parts[0].strip()
        title = parts[1].strip()
    elif '-' in name_without_ext:
        parts = name_without_ext.split('-', 1)
        artist = parts[0].strip()
        title = parts[1].strip()
    # Pattern 3: %26 for & in URL encoded
    elif '%26' in name_without_ext:
        decoded = unquote(name_without_ext)
        if '-' in decoded:
            parts = decoded.split('-', 1)
            artist = parts[0].strip()
            title = parts[1].strip()
    else:
        title = name_without_ext.strip()
    
    # Clean up percent encoding
    if artist:
        artist = unquote(artist)
    if title:
        title = unquote(title)
    if album:
        album = unquote(album)
    
    # Clean up weird characters
    artist = re.sub(r'\s+', ' ', artist).strip(' ????')
    title = re.sub(r'\s+', ' ', title).strip(' ????')
    album = re.sub(r'\s+', ' ', album).strip(' ????')
    
    # Replace empty or all ? marks with empty string
    if not artist or all(c == '?' for c in artist):
        artist = ''
    if not title or all(c == '?' for c in title):
        title = ''
    if not album or all(c == '?' for c in album):
        album = ''
    
    return album, artist, title, extension

def main():
    input_file = 'output.csv'
    output_file = 'output_completed.csv'
    
    with open(input_file, 'r', encoding='utf-8', errors='ignore') as infile:
        reader = csv.DictReader(infile)
        fieldnames = reader.fieldnames
        if not fieldnames:
            print("Error: No fieldnames found in CSV")
            return
        
        rows = []
        for row in reader:
            # Only try to fill if empty or question marks
            filename = row['file_name']
            
            # Parse from filename
            parsed_album, parsed_artist, parsed_title, _ = parse_filename(filename)
            
            # Update only if current is empty/questionable and we have something
            if (not row['album'] or all(c == '?' for c in row['album'])) and parsed_album:
                row['album'] = parsed_album
            if (not row['artist'] or all(c == '?' for c in row['artist'])) and parsed_artist:
                row['artist'] = parsed_artist
            if (not row['title'] or all(c == '?' for c in row['title'])) and parsed_title:
                row['title'] = parsed_title
            
            # Fix encoding issues
            for col in ['album', 'artist', 'title']:
                if row[col]:
                    # Check for garbled text and try to fix
                    try:
                        # If it's already okay, leave it
                        row[col].encode('utf-8').decode('utf-8')
                    except:
                        # Try to fix gbk encoding
                        try:
                            row[col] = row[col].encode('latin1').decode('gbk')
                        except:
                            try:
                                row[col] = row[col].encode('latin1').decode('utf-8')
                            except:
                                pass
            
            rows.append(row)
    
    with open(output_file, 'w', encoding='utf-8', newline='') as outfile:
        writer = csv.DictWriter(outfile, fieldnames=fieldnames)
        writer.writeheader()
        writer.writerows(rows)
    
    print(f"Completed! Output saved to {output_file}")

if __name__ == '__main__':
    main()


import csv
import sys

def count_empty_fields(filename, field):
    count = 0
    total = 0
    with open(filename, 'r', encoding='utf-8', errors='ignore') as f:
        reader = csv.DictReader(f)
        for row in reader:
            total += 1
            value = row[field]
            if not value or all(c == '?' for c in value):
                count += 1
    return count, total

filename = sys.argv[1] if len(sys.argv) > 1 else 'output_completed.csv'

original_artist_empty, total = count_empty_fields('output.csv', 'artist')
original_album_empty, _ = count_empty_fields('output.csv', 'album')
original_title_empty, _ = count_empty_fields('output.csv', 'title')

final_artist_empty, _ = count_empty_fields(filename, 'artist')
final_album_empty, _ = count_empty_fields(filename, 'album')
final_title_empty, _ = count_empty_fields(filename, 'title')

print(f"Total songs: {total}")
print()
print("Original:")
print(f"  Empty/questionable artist: {original_artist_empty}")
print(f"  Empty/questionable album: {original_album_empty}")  
print(f"  Empty/questionable title: {original_title_empty}")
print()
print(f"Final ({filename}):")
print(f"  Empty/questionable artist: {final_artist_empty} (filled {original_artist_empty - final_artist_empty})")
print(f"  Empty/questionable album: {final_album_empty} (filled {original_album_empty - final_album_empty})")
print(f"  Empty/questionable title: {final_title_empty} (filled {original_title_empty - final_title_empty})")

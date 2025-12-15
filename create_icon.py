"""
Create icon for Antigravity Cleaner
Generates a modern icon with PIL/Pillow
"""

try:
    from PIL import Image, ImageDraw, ImageFont
except ImportError:
    print("Installing Pillow...")
    import subprocess
    subprocess.check_call(['python', '-m', 'pip', 'install', 'Pillow'])
    from PIL import Image, ImageDraw, ImageFont

# Create a 512x512 image
size = 512
img = Image.new('RGBA', (size, size), (0, 0, 0, 0))
draw = ImageDraw.Draw(img)

# Background gradient (dark to darker)
for i in range(size):
    alpha = int(255 * (1 - i / size))
    draw.rectangle([(0, i), (size, i+1)], fill=(10, 10, 10, 255))

# Draw rounded rectangle background
margin = 40
draw.rounded_rectangle(
    [(margin, margin), (size-margin, size-margin)],
    radius=60,
    fill=(20, 20, 20, 255),
    outline=(0, 212, 255, 255),
    width=8
)

# Draw a stylized broom/rocket icon
center_x = size // 2
center_y = size // 2

# Rocket body (cyan)
rocket_points = [
    (center_x, center_y - 120),  # top
    (center_x - 40, center_y + 60),  # bottom left
    (center_x + 40, center_y + 60),  # bottom right
]
draw.polygon(rocket_points, fill=(0, 212, 255, 255))

# Rocket window
draw.ellipse(
    [(center_x - 25, center_y - 60), (center_x + 25, center_y - 10)],
    fill=(255, 0, 110, 255)
)

# Rocket flames (pink/magenta)
flame_points = [
    (center_x - 30, center_y + 60),
    (center_x - 20, center_y + 120),
    (center_x, center_y + 80),
]
draw.polygon(flame_points, fill=(255, 0, 110, 255))

flame_points2 = [
    (center_x + 30, center_y + 60),
    (center_x + 20, center_y + 120),
    (center_x, center_y + 80),
]
draw.polygon(flame_points2, fill=(255, 0, 110, 255))

# Center flame
flame_points3 = [
    (center_x - 10, center_y + 60),
    (center_x, center_y + 140),
    (center_x + 10, center_y + 60),
]
draw.polygon(flame_points3, fill=(255, 100, 150, 255))

# Add sparkles
sparkle_positions = [
    (center_x - 100, center_y - 80),
    (center_x + 100, center_y - 60),
    (center_x - 80, center_y + 20),
    (center_x + 90, center_y + 30),
]

for x, y in sparkle_positions:
    # Draw a star
    star_size = 15
    draw.line([(x, y - star_size), (x, y + star_size)], fill=(255, 255, 255, 255), width=3)
    draw.line([(x - star_size, y), (x + star_size, y)], fill=(255, 255, 255, 255), width=3)

# Save as PNG
img.save('icon.png', 'PNG')
print("Created icon.png (512x512)")

# Create ICO file (multiple sizes)
icon_sizes = [(256, 256), (128, 128), (64, 64), (48, 48), (32, 32), (16, 16)]
img.save('icon.ico', format='ICO', sizes=icon_sizes)
print("Created icon.ico (multi-size)")

# Create ICNS for macOS (if on macOS)
import platform
if platform.system() == 'Darwin':
    try:
        import subprocess
        # Create iconset directory
        subprocess.run(['mkdir', '-p', 'icon.iconset'])
        
        # Generate different sizes
        for size_val in [16, 32, 64, 128, 256, 512]:
            resized = img.resize((size_val, size_val), Image.Resampling.LANCZOS)
            resized.save(f'icon.iconset/icon_{size_val}x{size_val}.png')
            
            # Also create @2x versions
            if size_val <= 256:
                resized_2x = img.resize((size_val * 2, size_val * 2), Image.Resampling.LANCZOS)
                resized_2x.save(f'icon.iconset/icon_{size_val}x{size_val}@2x.png')
        
        # Convert to ICNS
        subprocess.run(['iconutil', '-c', 'icns', 'icon.iconset'])
        print("Created icon.icns (macOS)")
    except Exception as e:
        print(f"Could not create ICNS: {e}")

print("\nIcon generation complete!")
print("Files created:")
print("  - icon.png (512x512)")
print("  - icon.ico (Windows)")
if platform.system() == 'Darwin':
    print("  - icon.icns (macOS)")

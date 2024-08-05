from flask import Flask, request, jsonify
import requests

app = Flask(__name__)

mood_to_playlist = {
    "happy": [],
    "sad": [],
    "relaxed": [],
    "angry": []
}

@app.route('/update_likes', methods=['POST'])
def update_likes():
    data = request.json
    track_id = data.get('track_id')
    mood = data.get('mood', 'happy')
    
    mood_to_playlist[mood].append(track_id)
    
    return jsonify({"status": "success", "message": f"Track {track_id} added to {mood} playlist"})

@app.route('/get_playlist', methods=['POST'])
def get_playlist():
    data = request.json
    mood = data.get('mood', 'happy')  # По умолчанию используем "happy"
    
    playlist = mood_to_playlist.get(mood, [])
    
    return jsonify({"playlist": playlist})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)

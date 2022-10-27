import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'golang-live-chat';

  // Enable pusher logging - don't include this in production
    Pusher.logToConsole = true;

    pusher = new this.Pusher('51f3c1a8e5a382b22ce9', {
      cluster: 'eu'
    });

    channel = pusher.subscribe('my-channel');
    channel.bind('my-event', function(data) {
      alert(JSON.stringify(data));
    });
}

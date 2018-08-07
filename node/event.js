const EventEmitter = require( 'events' );
class MyEmitter extends EventEmitter {}

const myEmitter = new MyEmitter();

myEmitter.on( 'event', ( a, b ) => {
    myEmitter.emit( 'error', new Error( 'whoops!' ));
    console.log('something has happended! ' + a + ' ' + b );
});

myEmitter.on( 'error', ( err ) => {
    console.error( 'error' );
});

myEmitter.emit( 'event', 'a', 'b' );

var svg            = d3.select( "svg" ),
    topicmapholder = document.getElementById( "topicmapholder" ),
    width          = topicmapholder.clientWidth,
    height         = topicmapholder.clientHeight,
    regex1         = /topic(\d*)/,
    thisTopicNum   = regex1.exec( window.location.href )[ 1 ];

var basepath = window.location.href.split( '/' ).slice( 0, -5 ).join( '/' );

var simulation = d3.forceSimulation().nodes( visualizationData.nodes );

var link_force   = d3.forceLink( visualizationData.links ).id( function ( d ) {
    return d.id;
} );
//add forces
var charge_force = d3.forceManyBody().strength( -3500 ).distanceMax( 500 ).distanceMin( 100 );
// var charge_force = d3.forceManyBody().distanceMin(160);
simulation
    .force( "charge_force", charge_force )
    .force( "center_force", d3.forceCenter( width / 4, height / 3 ) )
    .force( "links", link_force );
var holdAll = svg.append( "g" ).attr( "class", "holdAll" );


var link = holdAll.append( "g" )
    .attr( "class", "links" )
    .selectAll( "line" )
    .data( visualizationData.links )
    .enter().append( "line" )
    .attr( "stroke-width", .7 )
    .style( "stroke", "black" );

var node = holdAll.append( "g" )
    .attr( "class", "nodes" )
    .selectAll( "circle" )
    .data( visualizationData.nodes )
    .enter().append( "g" )
    .attr( "class", "node" )
    .attr( "id", getId )
    .attr( "data-ocount", function ( d ) {
        return d.ocount;
    } )
    .attr( "data-nodenum", function ( d ) {
        return d.id;
    } )
    .attr( "class", isActive )
    .call( d3.drag()
               .on( "start", dragstarted )
               .on( "drag", dragged )
               .on( "end", dragended ) );

node.append( "circle" )
    .attr( "r", calculateRadius )
    .attr( "class", "circle" );
node.append( "text" )
    .attr( "dx", determineOffset )
    .attr( "dy", ".35em" )
    .text( function ( d ) {
        return d.name;
    } )
    .attr( "font-size", "1rem" );
node.on( "click", function ( d ) {
    window.location.href = basepath + '/' + d.path;
} );


function calculateRadius( d ) {
    return (
        7 + d.ocount
    );
}

function determineOffset( d ) {
    return (
        10 + d.ocount
    );
}

function getId( d ) {

    return (
        "nodenum" + d.id
    );
}

function getNodeNum( d ) {

    return (
        d.id
    );
}

function isActive( d ) {

    if ( d.id == thisTopicNum ) {
        return "node active"
    } else {
        return "node";
    }

}

function dragstarted( d ) {
    console.log( "dragstarted" );
    // simulation.restart();
    //  simulation.alpha(0.7);
    d.fx = d.x;
    d.fy = d.y;
}

function dragged( d ) {
    d.fx = d3.event.x;
    d.fy = d3.event.y;
}

function dragended( d ) {
    d.fx = null;
    d.fy = null;
    simulation.alphaTarget( 0.1 );
}

function tickActions() {
    //update circle positions to reflect node updates on each tick of the simulation
    /* node
         .attr("cx", function(d) { return d.x; })
         .attr("cy", function(d) { return d.y; });*/

    node.attr( "transform", function ( d ) {
        return "translate(" + d.x + "," + d.y + ")";
    } );
    //update link positions
    //simply tells one end of the line to follow one node around
    //and the other end of the line to follow the other node around
    link
        .attr( "x1", function ( d ) {
            return d.source.x;
        } )
        .attr( "y1", function ( d ) {
            return d.source.y;
        } )
        .attr( "x2", function ( d ) {
            return d.target.x;
        } )
        .attr( "y2", function ( d ) {
            return d.target.y;
        } );
}


//add zoom capabilities
var zoom_handler = d3.zoom().on( "zoom", zoom_actions );

zoom_handler( svg );

//Zoom functions
function zoom_actions() {
    holdAll.attr( "transform", d3.event.transform );
}

simulation.on( "tick", tickActions );
var renderedWidth  = d3.select( '.holdAll' ).node().getBoundingClientRect().width;
var renderedHeight = d3.select( '.holdAll' ).node().getBoundingClientRect().height;
console.log( "width of visualization " + renderedWidth );
console.log( "height of visualization " + renderedHeight );

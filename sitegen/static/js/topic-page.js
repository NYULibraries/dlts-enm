var basepath               = window.location.href.split( '/' ).slice( 0, -5 ).join( '/' ),

    activeTopicId          = parseInt( /(\d*).html/.exec( window.location.href )[ 1 ] ),

    svg                    = d3.select( "svg" ),

    forceSimulationElement = document.getElementById( "force-simulation" ),
    width                  = forceSimulationElement.clientWidth,
    height                 = forceSimulationElement.clientHeight,

    minNodeRadius = 7,
    minNodeOffset = 10,

    // visualizationData is defined and initialized in a previous <script> tag
    forceSimulation = d3.forceSimulation().nodes( visualizationData.nodes ),
    linkForce = d3.forceLink( visualizationData.links )
        .id(
            function ( d ) {
                return d.id;
            }
        ),
    chargeForce = d3.forceManyBody()
        .strength( -3500 )
        .distanceMax( 500 )
        .distanceMin( 100 ),

    wrapperG = svg.append( "g" ).attr( "class", "wrapper-g" ),

    link = wrapperG.append( "g" )
            .attr( "class", "links" )
        .selectAll( "line" )
        .data( visualizationData.links )
        .enter().append( "line" )
            .attr( "stroke-width", .7 )
            .style( "stroke", "black" ),

    node = wrapperG.append( "g" )
            .attr( "class", "nodes" )
        .selectAll( "circle" )
        .data( visualizationData.nodes )
        .enter().append( "g" )
            .attr( "class", "node" )
            .attr( "id", getNodeId )
            .attr( "data-ocount", function ( d ) {
                return d.ocount;
            } )
            .attr( "data-nodenum", function ( d ) {
                return d.id;
            } )
            .attr( "class", isActive )
            .call( d3.drag()
                .on( "start", dragStarted )
                .on( "drag", dragged )
                .on( "end", dragEnded ) ),

    zoomHandler = d3.zoom().on( "zoom", zoom_actions ),

    renderedWidth, renderedHeight;


forceSimulation.force( "charge_force", chargeForce )
    .force( "center_force", d3.forceCenter( width / 4, height / 3 ) )
    .force( "links", linkForce );

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

zoomHandler( svg );

// This must be called after enabling zoom handling, otherwise the graph nodes
// will no longer be draggable.  Not sure why this is.  Laura's original code
// made the calls in this order, and when they were reversed, dragging broke.
// This ticket might pertain:
// "Unable to drag nodes in a force layout graph when zooming is added #1412"
// https://github.com/d3/d3/issues/1412
forceSimulation.on( "tick", tickActions );

// Zoom functions
function zoom_actions() {
    wrapperG.attr( "transform", d3.event.transform );
}

function calculateRadius( d ) {
    return minNodeRadius + d.ocount;
}

function determineOffset( d ) {
    return minNodeOffset + d.ocount;
}

function getNodeId( d ) {

    return "nodenum" + d.id;
}

function isActive( d ) {

    if ( d.id === activeTopicId ) {
        return "node active";
    } else {
        return "node";
    }

}

function dragStarted( d ) {
    d.fx = d.x;
    d.fy = d.y;
}

function dragged( d ) {
    d.fx = d3.event.x;
    d.fy = d3.event.y;
}

function dragEnded( d ) {
    d.fx = null;
    d.fy = null;

    forceSimulation.alphaTarget( 0.1 );
}

function tickActions() {
    node.attr( "transform", function ( d ) {
        return "translate(" + d.x + "," + d.y + ")";
    } );
    // Update link positions
    // Simply tells one end of the line to follow one node around
    // and the other end of the line to follow the other node around
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

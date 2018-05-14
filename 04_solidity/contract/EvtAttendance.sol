pragma solidity ^0.4.23;


/**
 * @title EvtAttendance
 * @author Vincent Serpoul
 * @dev contract for managing attendance today
 */
contract EvtAttendance {

    Evt[] evts;

    struct Evt {
        uint256 id;
        address organizer;
        string title;
        address[] attendees;
        mapping(address => bool) hasAttended;
    }

    event EvtCreated(address _organizer, uint256 _id, string _title);
    event EvtAttended(address _who, uint256 _id);

    /**
    * @dev allows attendees to attend
    * @param _title string title of the event
    */
    function createEvt(string _title)
    public
    returns(uint256)
    {
        Evt memory evt;
        evt.id = evts.length;
        evt.organizer = msg.sender;
        evt.title = _title;
        evts.push(evt);

        emit EvtCreated(msg.sender, evt.id, _title);

        return evt.id;
    }

    /**
    * @dev allows attendees to attend
    * @param _evtId uint256 event id
    */
    function attend(uint256 _evtId)
    public
    {
        evts[_evtId].attendees.push(msg.sender);
        evts[_evtId].hasAttended[msg.sender] = true;

        emit EvtAttended(msg.sender, _evtId);
    }

    /**
    * @dev know if a user has attended
    * @param _evtId uint256 event id
    */
    function hasAttended(uint256 _evtId)
    view
    public
    returns(bool)
    {
        return evts[_evtId].hasAttended[msg.sender];
    }

    /**
    * @dev get an event
    * @param _evtId uint256 event id
    */
    function getEvent(uint256 _evtId)
    view
    public
    returns(uint256, address, string, address[]) {
        return (evts[_evtId].id, evts[_evtId].organizer, evts[_evtId].title, evts[_evtId].attendees);
    }

    /**
    * @dev get how many events have been created
    */
    function getEventCount()
    view
    public
    returns(uint256) {
        return evts.length;
    }
}
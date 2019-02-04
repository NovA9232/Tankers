function Cannon(parent)
  local c = {}
  c.parent = parent
  c.pos = parent.pos  -- Since it is a table, this will be a reference so no need to keep updating it.
  c.angle = 0
  c.w = 20
  c.h = 30

  function c:draw()
    lg.setColor(1, 1, 1, 1)
    lg.push()
      lg.rotate(self.angle+self.parent.angle-(math.pi/2))
      lg.rectangle("line", -(self.w/2), -(self.h/2), self.w, self.h)
      lg.rectangle("fill", -1.5, 0, 3, 30)
    lg.pop()
  end

  function c:update(dt)
    local mouseX, mouseY = love.mouse.getPosition()
    local angleToMouse = math.atan2((mouseY - self.pos.y), (mouseX - self.pos.x))
    --self.angle = angleToMouse
    local diff = self.angle-angleToMouse
    print(self.angle-angleToMouse)

    if diff < 0 or diff % (2*math.pi) > math.pi then
      print("egg")
      self.angle = self.angle + (TANK_CANNON_ROTATE_RATE*dt)+self.parent.angularVelocity
    else
      print("EGG")
      self.angle = self.angle - (TANK_CANNON_ROTATE_RATE*dt)+self.parent.angularVelocity
    end

    -- if math.abs(self.angle%math.pi)-angleToMouse < 0 then
    -- else
    --   self.angle = self.angle - (TANK_CANNON_ROTATE_RATE*dt)+self.parent.angularVelocity
    -- end
  end

  return c
end
